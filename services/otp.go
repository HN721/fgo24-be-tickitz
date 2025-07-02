package services

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateOTP() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	otp := r.Intn(900000) + 100000
	return strconv.Itoa(otp)
}
