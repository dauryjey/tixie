package handlers

import (
	"auth/db"
	"auth/models"
	"auth/utils"
	"encoding/json"
	"net/http"
)

type SignupPayload struct {
	models.User
	models.OptionalBase
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var payload SignupPayload

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := db.Connection()

	var user models.User

	where := models.User{
		Email: payload.Email,
	}

	isEmailUnique := db.First(&user, where)

	if isEmailUnique.Error != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Something wrong happened.")
		return
	}

	if isEmailUnique.Error == nil {
		utils.WriteErrorResponse(w, http.StatusConflict, "Email already exists.")
		return
	}
}
