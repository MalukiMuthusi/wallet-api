package models

type BasicError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
