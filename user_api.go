package goapi

import (
	"fmt"
	"net/http"
)

// Models

type User struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	Password        string     `json:"password,omitempty"`
	ProfileImageURL string     `json:"profile_image_url"`
	Profiles        Profiles   `json:"profiles"`
	DateFields      DateFields `json:"date_fields"`
}

type UserProfile struct {
	ID              string                 `json:"id"`
	UserID          string                 `json:"user_id"`
	TenantID        string                 `json:"tenant_id"`
	RoleID          string                 `json:"role_id"`
	IsProductStaff  bool                   `json:"is_product_staff"`
	IsActive        bool                   `json:"is_active"`
	AllowedProducts []Product              `json:"allowed_products"`
	PayTypes        []ProfilePayType       `json:"pay_types"`
	TimeSheets      []ProfileTimeSheet     `json:"time_sheets"`
	Reimbursements  []ProfileReimbursement `json:"reimbursements"`
	DateFields      DateFields             `json:"date_fields"`
}

type Profiles []UserProfile

type ProfilePayType struct {
	ID            string  `json:"id"`
	TenantID      string  `json:"tenant_id"`
	UserProfileID string  `json:"user_profile_id"`
	Name          string  `json:"name"`
	PayRate       float64 `json:"pay_rate"` // Use float64 for decimals
	EndDate       *string `json:"end_date"`
}

type ProfileTimeSheet struct {
	ID                string  `json:"id"`
	TenantID          string  `json:"tenant_id"`
	UserProfileID     string  `json:"user_profile_id"`
	UserName          string  `json:"user_name"`
	UserEmail         string  `json:"user_email"`
	PayType           string  `json:"pay_type"`
	PayRate           float64 `json:"pay_rate"`
	TimeIn            string  `json:"time_in"` // Use string to match time format
	ActualTimeIn      *string `json:"actual_time_in"`
	OrigTimeIn        *string `json:"orig_time_in"`
	LngIn             string  `json:"lng_in"`
	LatIn             string  `json:"lat_in"`
	ImageInURL        string  `json:"image_in_url"`
	TimeOut           *string `json:"time_out"`
	ActualTimeOut     *string `json:"actual_time_out"`
	OrigTimeOut       *string `json:"orig_time_out"`
	LngOut            string  `json:"lng_out"`
	LatOut            string  `json:"lat_out"`
	ImageOutURL       string  `json:"image_out_url"`
	Total             float64 `json:"total"`
	Note              string  `json:"note"`
	Exceptions        string  `json:"exceptions"`
	ExceptionsHandled bool    `json:"exceptions_handled"`
	ManuallyEntered   bool    `json:"manually_entered"`
	PayrollBatchID    *string `json:"payroll_batch_id"`
}

type ProfileReimbursement struct {
	ID             string  `json:"id"`
	TenantID       string  `json:"tenant_id"`
	UserProfileID  string  `json:"user_profile_id"`
	UserName       string  `json:"user_name"`
	UserEmail      string  `json:"user_email"`
	Date           string  `json:"date"` // Use string to match date format
	Amount         float64 `json:"amount"`
	Reason         string  `json:"reason"`
	ReceiptURL     string  `json:"receipt_url"`
	Status         string  `json:"status"`
	ApprovedAmount float64 `json:"approved_amount"`
	Note           string  `json:"note"`
	PayrollBatchID *string `json:"payroll_batch_id"`
}

type DateFields struct {
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}

// GetUsers retrieves a list of users
func GetUsers() ([]User, error) {
	var users []User
	response, err := makeRequest("GET", "/users", nil)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &users); err != nil {
		return nil, err
	}
	return users, nil
}

// GetUser retrieves a single user by ID
func GetUser(userID string) (*User, error) {
	var user User
	response, err := makeRequest("GET", "/users/"+userID, nil)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user
func CreateUser(user User) (*User, error) {
	var createdUser User
	response, err := makeRequest("POST", "/users", user)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &createdUser); err != nil {
		return nil, err
	}
	return &createdUser, nil
}

// DeleteUser deletes a user by ID
func DeleteUser(userID string) error {
	response, err := makeRequest("DELETE", "/users/"+userID, nil)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete user, status code: %d", response.StatusCode)
	}
	return nil
}

// UpdateUser updates an existing user
func UpdateUser(userID string, user User) (*User, error) {
	var updatedUser User
	response, err := makeRequest("PUT", "/users/"+userID, user)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &updatedUser); err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

// GetTimeSheets retrieves time sheets for a specific user profile
func GetTimeSheets(userProfileID string) ([]ProfileTimeSheet, error) {
	var timeSheets []ProfileTimeSheet
	response, err := makeRequest("GET", "/time_sheets?user_profile_id="+userProfileID, nil)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &timeSheets); err != nil {
		return nil, err
	}
	return timeSheets, nil
}

// CreateTimeSheet creates a new time sheet
func CreateTimeSheet(timeSheet ProfileTimeSheet) (*ProfileTimeSheet, error) {
	var createdTimeSheet ProfileTimeSheet
	response, err := makeRequest("POST", "/time_sheets", timeSheet)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &createdTimeSheet); err != nil {
		return nil, err
	}
	return &createdTimeSheet, nil
}

// UpdateTimeSheet updates an existing time sheet
func UpdateTimeSheet(timeSheetID string, timeSheet ProfileTimeSheet) (*ProfileTimeSheet, error) {
	var updatedTimeSheet ProfileTimeSheet
	response, err := makeRequest("PUT", "/time_sheets/"+timeSheetID, timeSheet)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &updatedTimeSheet); err != nil {
		return nil, err
	}
	return &updatedTimeSheet, nil
}

// DeleteTimeSheet deletes a time sheet by ID
func DeleteTimeSheet(timeSheetID string) error {
	response, err := makeRequest("DELETE", "/time_sheets/"+timeSheetID, nil)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete time sheet, status code: %d", response.StatusCode)
	}
	return nil
}

// GetReimbursements retrieves reimbursements for a specific user profile
func GetReimbursements(userProfileID string) ([]ProfileReimbursement, error) {
	var reimbursements []ProfileReimbursement
	response, err := makeRequest("GET", "/reimbursements?user_profile_id="+userProfileID, nil)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &reimbursements); err != nil {
		return nil, err
	}
	return reimbursements, nil
}

// CreateReimbursement creates a new reimbursement
func CreateReimbursement(reimbursement ProfileReimbursement) (*ProfileReimbursement, error) {
	var createdReimbursement ProfileReimbursement
	response, err := makeRequest("POST", "/reimbursements", reimbursement)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &createdReimbursement); err != nil {
		return nil, err
	}
	return &createdReimbursement, nil
}

// UpdateReimbursement updates an existing reimbursement
func UpdateReimbursement(reimbursementID string, reimbursement ProfileReimbursement) (*ProfileReimbursement, error) {
	var updatedReimbursement ProfileReimbursement
	response, err := makeRequest("PUT", "/reimbursements/"+reimbursementID, reimbursement)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &updatedReimbursement); err != nil {
		return nil, err
	}
	return &updatedReimbursement, nil
}

// DeleteReimbursement deletes a reimbursement by ID
func DeleteReimbursement(reimbursementID string) error {
	response, err := makeRequest("DELETE", "/reimbursements/"+reimbursementID, nil)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete reimbursement, status code: %d", response.StatusCode)
	}
	return nil
}

// ClockIn records a clock-in for a specific user profile
func ClockIn(userProfileID string, timeSheet ProfileTimeSheet) (*ProfileTimeSheet, error) {
	var clockedInTimeSheet ProfileTimeSheet
	response, err := makeRequest("POST", "/clock_in?user_profile_id="+userProfileID, timeSheet)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &clockedInTimeSheet); err != nil {
		return nil, err
	}
	return &clockedInTimeSheet, nil
}

// ClockOut records a clock-out for a specific user profile
func ClockOut(userProfileID string, timeSheet ProfileTimeSheet) (*ProfileTimeSheet, error) {
	var clockedOutTimeSheet ProfileTimeSheet
	response, err := makeRequest("POST", "/clock_out?user_profile_id="+userProfileID, timeSheet)
	if err != nil {
		return nil, err
	}
	if err := parseJSONResponse(response, &clockedOutTimeSheet); err != nil {
		return nil, err
	}
	return &clockedOutTimeSheet, nil
}
