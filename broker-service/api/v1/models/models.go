package models

type RequestProfile struct {
	Id       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
