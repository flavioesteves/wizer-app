package models

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type RequestUser struct {
	Id       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
