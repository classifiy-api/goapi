package main

import (
	"fmt"
	"log"
	"time"

	"goapi"
)

func handle() {
	// Set the base URL and any necessary configurations
	goapi.SetBaseURL("https://api.example.com")
	goapi.SetTenantName("default")

	// Example: Create a new product
	newProduct := goapi.Product{
		Name:        "Gymnastics Class",
		Description: "A class for young gymnasts.",
	}

	product, err := goapi.CreateProduct(newProduct)
	if err != nil {
		log.Fatalf("Error creating product: %v", err)
	}
	fmt.Printf("Created Product: %+v\n", product)

	// Example: Retrieve a product by ID
	productID := product.ID
	retrievedProduct, err := goapi.GetProduct(productID)
	if err != nil {
		log.Fatalf("Error retrieving product: %v", err)
	}
	fmt.Printf("Retrieved Product: %+v\n", retrievedProduct)

	// Example: Update a product
	updatedProduct := goapi.Product{
		ID:          productID,
		Name:        "Updated Gymnastics Class",
		Description: "Updated description.",
	}

	product, err = goapi.UpdateProduct(productID, updatedProduct)
	if err != nil {
		log.Fatalf("Error updating product: %v", err)
	}
	fmt.Printf("Updated Product: %+v\n", product)

	// Example: Delete a product
	err = goapi.DeleteProduct(productID)
	if err != nil {
		log.Fatalf("Error deleting product: %v", err)
	}
	fmt.Println("Product deleted successfully.")

	// Example: Create a product schedule
	productSchedule := goapi.ProductSchedule{
		TenantID:  "tenant123",
		ProductID: "product123",
		BeginDate: time.Now().Format("2006-01-02"),
		EndDate:   time.Now().AddDate(0, 1, 0).Format("2006-01-02"),
		TimeZone:  "America/New_York",
		Sessions: []goapi.ProductScheduleSession{
			{
				ID:                "session123",
				TenantID:          "tenant123",
				ProductScheduleID: "schedule123",
				LocationID:        "location123",
				Day:               "Monday",
				BeginTime:         "09:00",
				DurationMinutes:   60,
				Staff: []goapi.ProductScheduleSessionUser{
					{
						TenantID:                 "tenant123",
						ProductScheduleSessionID: "session123",
						UserID:                   "user123",
						User:                     &goapi.User{ID: "user123", Name: "John Doe"},
						Role:                     "Instructor",
					},
				},
				Resources: []goapi.ProductScheduleSessionResource{
					{
						TenantID:                 "tenant123",
						ProductScheduleSessionID: "session123",
						ResourceID:               "resource123",
						BeginTime:                "09:00",
						DurationMinutes:          60,
					},
				},
			},
		},
	}

	schedule, err := goapi.CreateProductSchedule(productSchedule)
	if err != nil {
		log.Fatalf("Error creating product schedule: %v", err)
	}
	fmt.Printf("Created Product Schedule: %+v\n", schedule)

	// Example: Retrieve a product schedule by ID
	scheduleID := schedule.ID
	retrievedSchedule, err := goapi.GetProductSchedule(scheduleID)
	if err != nil {
		log.Fatalf("Error retrieving product schedule: %v", err)
	}
	fmt.Printf("Retrieved Product Schedule: %+v\n", retrievedSchedule)

	// Example: Create a product schedule session
	session := goapi.ProductScheduleSession{
		ID:                "session123",
		TenantID:          "tenant123",
		ProductScheduleID: "schedule123",
		LocationID:        "location123",
		Day:               "Monday",
		BeginTime:         "09:00",
		DurationMinutes:   60,
		Staff: []goapi.ProductScheduleSessionUser{
			{
				TenantID:                 "tenant123",
				ProductScheduleSessionID: "session123",
				UserID:                   "user123",
				User:                     &goapi.User{ID: "user123", Name: "John Doe"},
				Role:                     "Instructor",
			},
		},
		Resources: []goapi.ProductScheduleSessionResource{
			{
				TenantID:                 "tenant123",
				ProductScheduleSessionID: "session123",
				ResourceID:               "resource123",
				BeginTime:                "09:00",
				DurationMinutes:          60,
			},
		},
	}

	newSession, err := goapi.CreateProductScheduleSession(session)
	if err != nil {
		log.Fatalf("Error creating product schedule session: %v", err)
	}
	fmt.Printf("Created Product Schedule Session: %+v\n", newSession)

	// Example: Retrieve a product schedule session by ID
	sessionID := newSession.ID
	retrievedSession, err := goapi.GetProductScheduleSession(sessionID)
	if err != nil {
		log.Fatalf("Error retrieving product schedule session: %v", err)
	}
	fmt.Printf("Retrieved Product Schedule Session: %+v\n", retrievedSession)
}
