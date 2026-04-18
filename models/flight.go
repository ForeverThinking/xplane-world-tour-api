package models

type Flight struct {
	FlightID        int    `json:"flightId"`
	OriginICAO      string `json:"originIcao"`
	DestinationICAO string `json:"destinationIcao"`
	AirframeID      int    `json:"airframeId"`
}

