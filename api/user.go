package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ijaybaihaqi/heli-api/model"
)

func (api *API) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit /auth/register")

	var creds model.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	if creds.Username == "" || creds.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Bad Request"})
		return
	}

	if api.userService.CheckPassLength(creds.Password) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Please provide a password of more than 5 characters"})
		return
	}

	if api.userService.CheckPassAlphabet(creds.Password) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Please use Password with Contains non Alphabetic Characters"})
		return
	}

	// Hash password
	hashedPass, err := api.userService.HashPassword(creds.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	creds.Password = hashedPass

	err = api.userService.Register(creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "User Registered"})
}

func (api *API) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit /auth/login")

	var creds model.User

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	if creds.Username == "" || creds.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Bad Request"})
		return
	}

	if api.userService.CheckPassLength(creds.Password) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Please provide a password of more than 5 characters"})
		return
	}

	if api.userService.CheckPassAlphabet(creds.Password) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Please use Password with Contains non Alphabetic Characters"})
		return
	}

	user, err := api.userService.Login(creds)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	expiresAt := time.Now().Add(time.Hour * 12)

	sessionToken, err := api.sessionService.GenerateJWT(user, expiresAt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	session := model.Session{Token: sessionToken, Username: creds.Username, Expiry: expiresAt}

	err = api.sessionService.SessionAvailName(session.Username)
	if err != nil {
		err = api.sessionService.AddSession(session)
	} else {
		err = api.sessionService.UpdateSession(session)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.LoginResponse{
		SessionToken: sessionToken,
	})
}

func (api *API) AutoLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit /auth/autologin")

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	// Extract the token from the "Bearer" scheme
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := api.sessionService.ValidateJWT(tokenString)
	if err != nil {
		http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	if username, ok := claims["username"].(string); ok {
		err = api.sessionService.SessionAvailName(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Token is not valid"})
			return
		}
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Message: "Welcome back",
	})
}

func (api *API) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit /auth/logout")

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	// Extract the token from the "Bearer" scheme
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	_, err := api.sessionService.ValidateJWT(tokenString)
	if err != nil {
		http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	api.sessionService.DeleteSession(tokenString)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "Logout Success"})
}
