package models

type Product struct {
	Id             string  `json:"id"`
	Width          int     `json:"width"`
	Height         int     `json:"height"`
	Length         int     `json:"length"`
	Weight         float32 `json:"weight"`
	InsuranceValue float32 `json:"insurance_value"`
	Quantity       int     `json:"quantity"`
}

type Package struct {
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Length int     `json:"length"`
	Weight float32 `json:"weight"`
}

type ShipmentOptions struct {
	Receipt        bool    `json:"receipt"`
	OwnHand        bool    `json:"own_hand"`
	InsuranceValue float32 `json:"insurance_value"`
}
