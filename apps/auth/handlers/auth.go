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

type signupResponse struct {
	Token string `json:"token"`
}

type loginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var payload signupPayload

	payloadErr := json.NewDecoder(r.Body).Decode(&payload)

	if payloadErr != nil {
		http.Error(w, payloadErr.Error(), http.StatusBadRequest)
		return
	}

	db, dbErr := db.Connection()

	if dbErr != nil {
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

	jwt, jwtErr := utils.GenerateJWT(jwtPayload)

	if jwtErr != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Something wrong happened.")
		return
	}

	response := signupResponse{Token: fmt.Sprintf("Bearer %s", *jwt)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var payload loginPayload

	payloadErr := json.NewDecoder(r.Body).Decode(&payload)

	if payloadErr != nil {
		http.Error(w, payloadErr.Error(), http.StatusBadRequest)
		return
	}

	var foundUser models.User

	where := models.User{
		Email: payload.Email,
	}

	db, dbErr := db.Connection()

	if dbErr != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Something wrong happened.")
		return
	}

	queryResult := db.Limit(1).Find(&foundUser, where)

	if queryResult.Error != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Something wrong happened.")
		return
	}

	if queryResult.RowsAffected == 0 {
		utils.WriteErrorResponse(w, http.StatusNotFound, "User not found.")
		return
	}

	comparePasswordErr := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(payload.Password))

	if comparePasswordErr != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "Your password is wrong.")
		return
	}

	jwtPayload := utils.JWTPayload{
		Email:  payload.Email,
		UserID: foundUser.ID,
	}

	jwt, jwtErr := utils.GenerateJWT(jwtPayload)

	if jwtErr != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Something wrong happened.")
		return
	}

	response := loginResponse{Token: fmt.Sprintf("Bearer %s", *jwt)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
