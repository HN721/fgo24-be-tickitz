package utils

type Response struct {
	Success bool   `json:"success" form:"sucess"`
	Message string `json:"message" form:"message"`
	Results any    `json:"results,omitempty" form:"results"`
	Error   string `json:"error,omitempty" `
	Token   string `json:"token,omitempty"`
}
