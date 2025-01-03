package model

type UserApplianceRequest struct {
	ID          string `json:"id"`
	ApplianceID uint   `json:"appliance_id"`
	Room        string `json:"room"`
}
