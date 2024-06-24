package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"xml-reader-api/internal/entity"

	"github.com/go-chi/jwtauth"
)

func TestLoginHandler(t *testing.T) {
	jwtAuth := jwtauth.New("HS256", []byte("secret"), nil)

	mockRepo := &MockUserRepository{
		GetUserByEmailFn: func(email string) (*entity.User, error) {
			if email == "james@foo.com" {
				return &entity.User{
					ID:       1,
					Name:     "James Foo",
					Email:    "james@foo.com",
					Password: "$2a$10$LL18McIxUibXoclRUR408ultpBzRZOmFN52BzneqWWECZwxxQfhuG", // abc123
				}, nil
			}
			return nil, errors.New("user not found")
		},
	}

	handler := NewAuthHandler(mockRepo)

	tests := []struct {
		name             string
		input            JwtInputDto
		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "Successfull login",
			input: JwtInputDto{
				Email:    "james@foo.com",
				Password: "abc123",
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"message":"User logged in successfully","error":false,"data":{"id":1,"email":"james@foo.com","name":"James Foo"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.input)
			req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
			if err != nil {
				t.Fatal(err)
			}

			ctx := context.WithValue(req.Context(), "jwt", jwtAuth)
			ctx = context.WithValue(ctx, "jwt-exp", 1)
			req = req.WithContext(ctx)

			rr := httptest.NewRecorder()
			handler.LoginHandler(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}
		})
	}
}
