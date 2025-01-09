package model

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	SessionToken string `json:"session_token"`
}

type ApplianceResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserApplianceResponse struct {
	ID            uint    `json:"id"`
	Appliance     string  `json:"appliance"`
	Image         string  `json:"image"`
	Energy        float64 `json:"energy"`
	Room          string  `json:"room"`
	CurrentStatus string  `json:"current_status"`
}

type UserApplianceRoomResponse struct {
	Room string `json:"room"`
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

type CurrentStatusResponse struct {
	Status string `json:"status"`
}

type TapasResponse struct {
	Answer      string   `json:"answer"`
	Coordinates [][]int  `json:"coordinates"`
	Cells       []string `json:"cells"`
	Aggregator  string   `json:"aggregator"`
}

type ChatResponse struct {
	GeneratedText string `json:"generated_text"`
}
