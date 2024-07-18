package main

import (
	"fmt"
	"log"

	"github.com/classify-api/goapi"
)

func main() {
	goapi.SetBaseURL("https://dev-api.jumpupgym.com")
	goapi.SetTenantName("tenant1")

	err := goapi.Login("user2@test.com", "test")
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	jwt, err := goapi.GetJWT()
	if err != nil {
		log.Fatalf("Failed to get JWT: %v", err)
	}

	fmt.Printf("JWT Token: %s\n", jwt)

	// response, err := goapi.MakeRequest("GET", "/your-endpoint", nil)
	// if err != nil {
	// 	log.Fatalf("Request failed: %v", err)
	// }
	// defer response.Body.Close()

	// fmt.Printf("Response status: %s\n", response.Status)
}
