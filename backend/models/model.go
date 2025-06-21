package models

type Lead struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email,omitempty"`
	Time  string `json:"time"`
}