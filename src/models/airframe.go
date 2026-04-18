package models

type Airframe struct {
	AirframeID   int    `json:"airframeId"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
}
