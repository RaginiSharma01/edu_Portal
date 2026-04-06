package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateOTP() string {
	otp := rng.Intn(900000) + 100000 // 6 digits
	return fmt.Sprintf("%d", otp)
}
