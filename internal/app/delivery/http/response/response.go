package response

type ErrorResponse struct {
	Message string `json:"message"`
}

type AuthResponse struct {
	Token string `json:"token"`
}