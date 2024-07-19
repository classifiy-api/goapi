package goapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

var (
	BASE_URL   string
	jwtToken   string // Updated to store single JWT token
	TenantName string // New global variable for the tenant name
	client     = &http.Client{}
	mu         sync.Mutex // Mutex for thread-safe access to jwtToken
)

type LoginRequest struct {
	User User `json:"user"`
}

type loginResponse struct {
	Tokens map[string]string `json:"tokens"` // Updated to handle multiple tokens
}

func SetBaseURL(url string) {
	BASE_URL = url
}

func SetTenantName(tenant string) {
	TenantName = tenant
}

func Login(username, password string) error {
	url := fmt.Sprintf("%s/login", BASE_URL)

	user := User{
		Email:    username,
		Password: password,
	}

	loginRequest := LoginRequest{
		User: user,
	}

	requestBody, err := json.Marshal(loginRequest)
	if err != nil {
		return fmt.Errorf("error marshaling login request: %v", err)
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("error making POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("login request failed with status: %v", resp.Status)
	}

	var response loginResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("error decoding login response: %v", err)
	}

	// Pick the JWT token for the specified tenant from the tokens map
	token, ok := response.Tokens[TenantName]
	if !ok {
		return fmt.Errorf("token for tenant %s not found in response", TenantName)
	}

	mu.Lock()
	jwtToken = token // Store the selected token globally
	mu.Unlock()

	return nil
}

func GetJWT() (string, error) {
	mu.Lock()
	defer mu.Unlock()
	if jwtToken == "" {
		return "", fmt.Errorf("JWT token not set")
	}
	return jwtToken, nil
}

func makeRequest(method, path string, body interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", BASE_URL, path)

	var requestBody []byte
	var err error

	if body != nil {
		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %v", err)
		}
	} else {
		requestBody = []byte{}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating new request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	mu.Lock()
	if jwtToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwtToken))
	}
	mu.Unlock()

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	return resp, nil
}

// Helper function to parse JSON responses
func parseJSONResponse(response *http.Response, result interface{}) error {
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, result)
}

// Helper function to format query parameters
func formatQueryParams(params map[string]string) string {
	query := url.Values{}
	for key, value := range params {
		query.Add(key, value)
	}
	return query.Encode()
}
