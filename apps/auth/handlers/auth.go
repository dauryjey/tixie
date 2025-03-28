package handlers

import (
	"auth/db"
	"auth/models"
	"auth/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type signupPayload struct {
	models.User
	models.OptionalBase
}

type SignupResponse struct {
	Token string `json:"token"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var payload signupPayload

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := db.Connection()

	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Something wrong happened.")
		return
	}

	var user models.User

	where := models.User{
		Email: payload.Email,
	}

	isEmailUnique := db.Limit(1).Find(&user, where)

	if isEmailUnique.Error != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Something wrong happened.")
		return
	}

	if isEmailUnique.RowsAffected > 0 {
		utils.WriteErrorResponse(w, http.StatusConflict, "Email already exists.")
		return
	}

	encryptedPassword, hashError := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if hashError != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Something wrong happened.")
		return
	}

	newUser := models.User{
		Name:        payload.Name,
		Email:       payload.Email,
		IdentityDoc: payload.IdentityDoc,
		Password:    string(encryptedPassword),
	}

	createUserRes := db.Create(&newUser)

	if createUserRes.Error != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Something wrong happened.")
		return
	}

	jwtPayload := utils.JWTPayload{
		Email:  payload.Email,
		UserID: newUser.ID,
	}

	response := SignupResponse{Token: fmt.Sprintf("Bearer %s", utils.GenerateJWT(jwtPayload))}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
