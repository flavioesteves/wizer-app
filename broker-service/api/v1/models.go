package v1

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayLoad `json:"auth, omitempty"`
}

type AuthPayLoad struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JSONresponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data, omitempty"`
}
