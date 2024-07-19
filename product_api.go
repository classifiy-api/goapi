package goapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Models

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductSchedule struct {
	ID        string                   `json:"id"`
	TenantID  string                   `json:"tenant_id"`
	ProductID string                   `json:"product_id"`
	BeginDate string                   `json:"begin_date"` // ISO format date
	EndDate   string                   `json:"end_date"`   // ISO format date
	TimeZone  string                   `json:"time_zone"`  // IANA Time Zone
	Sessions  []ProductScheduleSession `json:"sessions"`
}

type ProductScheduleSession struct {
	ID                string                           `json:"id"`
	TenantID          string                           `json:"tenant_id"`
	ProductScheduleID string                           `json:"product_schedule_id"`
	LocationID        string                           `json:"location_id"`
	Day               string                           `json:"day"`        // Day of the week
	BeginTime         string                           `json:"begin_time"` // Format "HH:MM"
	DurationMinutes   int                              `json:"duration_minutes"`
	Staff             []ProductScheduleSessionUser     `json:"staff"`
	Resources         []ProductScheduleSessionResource `json:"resources"`
}

type ProductScheduleSessionUser struct {
	TenantID                 string `json:"tenant_id"`
	ProductScheduleSessionID string `json:"product_schedule_session_id"`
	UserID                   string `json:"user_id"`
	Role                     string `json:"role"`
}

type ProductScheduleSessionResource struct {
	TenantID                 string `json:"tenant_id"`
	ProductScheduleSessionID string `json:"product_schedule_session_id"`
	ResourceID               string `json:"resource_id"`
	BeginTime                string `json:"begin_time"` // Format "HH:MM"
	DurationMinutes          int    `json:"duration_minutes"`
}

type ProductScheduleSessionInstance struct {
	ID                       string                                     `json:"id"`
	TenantID                 string                                     `json:"tenant_id"`
	ProductID                string                                     `json:"product_id"`
	ProductScheduleID        string                                     `json:"product_schedule_id"`
	ProductScheduleSessionID string                                     `json:"product_schedule_session_id"`
	Date                     string                                     `json:"date"`       // ISO format date
	BeginTime                string                                     `json:"begin_time"` // Format "HH:MM"
	AttendanceRecords        []ProductScheduleSessionInstanceAttendance `json:"attendance_records"`
	Trials                   []ProductScheduleSessionInstanceTrials     `json:"trials"`
}

type ProductScheduleSessionInstanceAttendance struct {
	TenantID          string `json:"tenant_id"`
	SessionInstanceID string `json:"session_instance_id"`
	SubscriberID      string `json:"subscriber_id"`
	Attendance        string `json:"attendance"`
}

type ProductScheduleSessionInstanceTrials struct {
	TenantID          string `json:"tenant_id"`
	SessionInstanceID string `json:"session_instance_id"`
	SubscriberID      string `json:"subscriber_id"`
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

// GetProducts retrieves a list of products
func GetProducts() ([]Product, error) {
	var products []Product
	response, err := makeRequest("GET", "/products", nil)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &products); err != nil {
		return nil, err
	}
	return products, nil
}

// FilterProducts retrieves a list of products with the specified filters
func FilterProducts(filters map[string]string) ([]Product, error) {
	var products []Product
	query := "?" + formatQueryParams(filters)
	response, err := makeRequest("GET", "/products/filter"+query, nil)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &products); err != nil {
		return nil, err
	}
	return products, nil
}

// CreateProduct creates a new product
func CreateProduct(product Product) (*Product, error) {
	var createdProduct Product
	response, err := makeRequest("POST", "/products", product)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &createdProduct); err != nil {
		return nil, err
	}
	return &createdProduct, nil
}

// GetProduct retrieves a single product by ID
func GetProduct(productID string) (*Product, error) {
	var product Product
	response, err := makeRequest("GET", "/products/"+productID, nil)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &product); err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateProduct updates an existing product
func UpdateProduct(productID string, product Product) (*Product, error) {
	var updatedProduct Product
	response, err := makeRequest("PUT", "/products/"+productID, product)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &updatedProduct); err != nil {
		return nil, err
	}
	return &updatedProduct, nil
}

// DeleteProduct deletes a product by ID
func DeleteProduct(productID string) error {
	response, err := makeRequest("DELETE", "/products/"+productID, nil)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete product, status code: %d", response.StatusCode)
	}
	return nil
}

// CreateProductSchedule creates a new product schedule
func CreateProductSchedule(schedule ProductSchedule) (*ProductSchedule, error) {
	var createdSchedule ProductSchedule
	response, err := makeRequest("POST", "/product_schedules", schedule)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &createdSchedule); err != nil {
		return nil, err
	}
	return &createdSchedule, nil
}

// GetProductSchedule retrieves a product schedule by ID
func GetProductSchedule(scheduleID string) (*ProductSchedule, error) {
	var schedule ProductSchedule
	response, err := makeRequest("GET", "/product_schedules/"+scheduleID, nil)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &schedule); err != nil {
		return nil, err
	}
	return &schedule, nil
}

// CreateProductScheduleSession creates a new product schedule session
func CreateProductScheduleSession(session ProductScheduleSession) (*ProductScheduleSession, error) {
	var createdSession ProductScheduleSession
	response, err := makeRequest("POST", "/product_schedule_sessions", session)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &createdSession); err != nil {
		return nil, err
	}
	return &createdSession, nil
}

// GetProductScheduleSession retrieves a product schedule session by ID
func GetProductScheduleSession(sessionID string) (*ProductScheduleSession, error) {
	var session ProductScheduleSession
	response, err := makeRequest("GET", "/product_schedule_sessions/"+sessionID, nil)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &session); err != nil {
		return nil, err
	}
	return &session, nil
}

// CreateProductScheduleSessionUser creates a new product schedule session user
func CreateProductScheduleSessionUser(sessionUser ProductScheduleSessionUser) (*ProductScheduleSessionUser, error) {
	var createdSessionUser ProductScheduleSessionUser
	response, err := makeRequest("POST", "/product_schedule_session_users", sessionUser)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &createdSessionUser); err != nil {
		return nil, err
	}
	return &createdSessionUser, nil
}

// CreateProductScheduleSessionResource creates a new product schedule session resource
func CreateProductScheduleSessionResource(sessionResource ProductScheduleSessionResource) (*ProductScheduleSessionResource, error) {
	var createdSessionResource ProductScheduleSessionResource
	response, err := makeRequest("POST", "/product_schedule_session_resources", sessionResource)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &createdSessionResource); err != nil {
		return nil, err
	}
	return &createdSessionResource, nil
}
