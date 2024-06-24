package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"xml-reader-api/internal/entity"
)

// MockUserRepository is a mock type for the user repository
type MockUserRepository struct {
	CreateUserFn     func(name, email, password string) (int64, error)
	GetUserByEmailFn func(email string) (*entity.User, error)
}

// CreateUser is a mock function for the CreateUser method
func (m *MockUserRepository) CreateUser(name, email, password string) (int64, error) {
	return m.CreateUserFn(name, email, password)
}

// GetUserByEmail is a mock function for the GetUserByEmail method
func (m *MockUserRepository) GetUserByEmail(email string) (*entity.User, error) {
	return m.GetUserByEmailFn(email)
}

func TestCreateUserHandler(t *testing.T) {
	mockRepo := &MockUserRepository{
		CreateUserFn: func(name, email, password string) (int64, error) {
			if email == "existing@user.com" {
				return 0, errors.New("UNIQUE constraint failed")
			}
			return 1, nil
		},
	}

	handler := NewUserHandler(mockRepo)

	tests := []struct {
		name             string
		input            CreateUserDto
		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "Successfull user creation",
			input: CreateUserDto{
				Name:     "James Foo",
				Email:    "james@foo.com",
				Password: "1234",
			},
			expectedStatus:   http.StatusCreated,
			expectedResponse: `{"message":"User created successfully","error":false,"data":{"id":1,"name":"James Foo","email":"james@foo.com"}}`,
		},
		{
			name: "Email Already Exists",
			input: CreateUserDto{
				Name:     "Jhon Doe",
				Email:    "existing@user.com",
				Password: "1234",
			},
			expectedStatus:   http.StatusConflict,
			expectedResponse: `{"message":"email already exists","error":true}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.input)
			req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			rr := httptest.NewRecorder()
			handler.CreateUserHandler(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, status)
			}
			if !strings.Contains(rr.Body.String(), tt.expectedResponse) {
				t.Errorf("Expected response to contain %s, got %s", tt.expectedResponse, rr.Body.String())
			}
		})
	}
}
