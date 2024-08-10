package models

type RequestUser struct {
	Id       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
