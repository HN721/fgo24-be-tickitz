package dto

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RegisterResquest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
type ChangePassword struct {
	Email       string `json:"email"`
	OTP         int    `json:"otp"`
	NewPassword string `json:"newpassword"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}
