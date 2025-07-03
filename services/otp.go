package services

import (
	"math/rand"
	"time"
)

func GenerateOTP() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	otp := r.Intn(900000) + 100000
	return otp
}
