package models

type Lead struct {
	FullName    string `json:"fullName"`
	Position    string `json:"position"`
	Company     string `json:"company"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Course_description     string `json:"message"`
	Time        string `json:"time"`
}