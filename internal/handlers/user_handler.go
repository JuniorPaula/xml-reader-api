package handlers

import (
	"encoding/json"
	"net/http"
	"xml-reader-api/internal/entity"
	"xml-reader-api/internal/repository"
)

type UserHandler struct {
	UserDB repository.UserInterface
}

func NewUserHandler(userDB repository.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Payload struct {
	Message string      `json:"message"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data"`
}

// CreateUser is a function that creates a new user
func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		Payload := Payload{
			Message: "Failed to decode request body",
			Error:   true,
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Payload)
		return
	}

	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		Payload := Payload{
			Message: "Failed to generate user",
			Error:   true,
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Payload)
		return
	}

	ltsID, err := h.UserDB.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		Payload := Payload{
			Message: "Failed to create user",
			Error:   true,
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Payload)
		return
	}
	user.ID = ltsID

	Payload := Payload{
		Message: "User created successfully",
		Data:    user,
		Error:   false,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Payload)
}
