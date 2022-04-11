package models

type Balance struct {
	Value string `json:"value,omitempty"`
	// Only one currency type is currently supported, euro
	Currency string `json:"currency,omitempty"`
	Date     string `json:"date"`
}
