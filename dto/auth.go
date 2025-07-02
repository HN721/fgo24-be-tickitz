package dto

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RegisterResquest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm-password"`
}
type ChangePassword struct {
	OldPassword string `json:"old-password"`
	NewPassword string `json:"new-password"`
}
