package dto

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RegisterResquest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
