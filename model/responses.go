package model

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type UserResponse struct {
	Username     string `json:"username"`
	SessionToken string `json:"session_token"`
	ExpiresAt    string `json:"expires_at"`
}

type ApplianceResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type UserApplianceResponse struct {
	ID        string `json:"id"`
	Appliance string `json:"appliance"`
	Room      string `json:"room"`
}

type ConsumptionResponse struct {
	ID         string    `json:"id"`
	ConsumedAt time.Time `json:"consumed_at"`
	Appliance  string    `json:"appliance"`
	Room       string    `json:"room"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
}

type ConsumptionIDResponse struct {
	ID string `json:"id"`
}
