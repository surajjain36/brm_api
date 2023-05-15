package otpgen

import (
	"math/rand"
	"strings"
	"time"
)

//GenerateOTP for OTP
func GenerateOTP(length int) string {
	rand.Seed(time.Now().Unix())
	var output strings.Builder

	//Lowercase, Uppercase and Number
	charSet := "0123456789"
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}

//OpacGen for OTP
func OpacGen(length int) string {
	rand.Seed(time.Now().Unix())
	var output strings.Builder

	//Lowercase, Uppercase and Number
	charSet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+;,."
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}
