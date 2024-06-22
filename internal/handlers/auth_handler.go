package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	"xml-reader-api/internal/repository"
	"xml-reader-api/internal/utils"

	"github.com/go-chi/jwtauth"
)

type AuthHandler struct {
	UserDB repository.UserInterface
}

func NewAuthHandler(userDB repository.UserInterface) *AuthHandler {
	return &AuthHandler{
		UserDB: userDB,
	}
}

type JwtInputDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginHandler is a function that logs in a user
func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// get the jwt token with context
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExp := r.Context().Value("jwt-exp").(int)

	var credentials JwtInputDto
	err := utils.ReadJSON(w, r, &credentials)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusUnprocessableEntity)
		return
	}
	// get user by email
	user, err := h.UserDB.GetUserByEmail(credentials.Email)
	if err != nil {
		utils.ErrorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}
	// validate password
	if !user.ValidatePassword(credentials.Password) {
		utils.ErrorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}
	fmt.Println("exp", jwtExp)
	// generate jwt token
	_, token, err := jwt.Encode(map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * time.Duration(jwtExp)).Unix(),
	})
	if err != nil {
		fmt.Println("error generating token", err)
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	data := struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		AccessToken string `json:"access_token"`
	}{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: token,
	}

	payload := utils.JsonResponse{
		Error:   false,
		Message: "Login successful",
		Data:    data,
	}
	utils.WriteJSON(w, http.StatusOK, payload)
}
