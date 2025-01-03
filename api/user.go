package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
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

	err = api.userService.Login(creds)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(5 * time.Hour)
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

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Path:    "/",
		Value:   sessionToken,
		Expires: expiresAt,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.UserResponse{
		Username:     creds.Username,
		SessionToken: sessionToken,
		ExpiresAt:    expiresAt.Format(time.RFC3339),
	})
}

func (api *API) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit /auth/logout")

	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	sessionToken := c.Value

	api.sessionService.DeleteSession(sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Path:    "/",
		Value:   "",
		Expires: time.Now(),
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "Logout Success"})
}
