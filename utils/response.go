package utils

type Response struct {
	Success bool   `json:"success" form:"sucess"`
	Message string `json:"message" form:"message"`
	Results any    `json:"results" form:"results"`
	Error   string `json:"error"`
}
