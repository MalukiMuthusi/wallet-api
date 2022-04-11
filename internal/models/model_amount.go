package models

type Amount struct {
	Value string `json:"value,omitempty" form:"value"`
	// Only one currency type is currently supported, euro
	Currency string `json:"currency,omitempty" form:"currency"`
}
