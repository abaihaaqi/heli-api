package model

type UserApplianceRequest struct {
	ID          uint   `json:"id"`
	ApplianceID uint   `json:"appliance_id"`
	Room        string `json:"room"`
}

type AnalyzeDataRequest struct {
	Table map[string][]string `json:"table"`
	Query string              `json:"query"`
}

type ChatRequest struct {
	Query string `json:"query"`
}
