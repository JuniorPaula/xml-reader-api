package handlers

import (
	"errors"
	"net/http"
	"strings"
	"xml-reader-api/internal/entity"
	"xml-reader-api/internal/repository"
	"xml-reader-api/internal/utils"
)

type UserHandler struct {
	UserDB repository.UserInterface
}

func NewUserHandler(userDB repository.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

type CreateUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUser is a function that creates a new user
func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input CreateUserDto
	err := utils.ReadJSON(w, r, &input)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusUnprocessableEntity)
		return
	}

	ltsID, err := h.UserDB.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			utils.ErrorJSON(w, errors.New("email already exists"), http.StatusConflict)
			return
		}
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	user.ID = ltsID
	payload := utils.JsonResponse{
		Message: "User created successfully",
		Data:    user,
	}
	utils.WriteJSON(w, http.StatusCreated, payload)
}
