package common

import (
	"log"
	"math/rand"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func PaisaToRupee(amount float64) float64 {
	return amount / 100
}
func RupeeToPaisa(amount float64) float64 {
	return amount * 100
}

func PaisaToRupeeStr(amount float64) string {
	return strconv.FormatFloat(amount/100, 'f', 2, 64)
}
func RupeeToPaisaStr(amount float64) string {
	return strconv.FormatFloat(amount*100, 'f', 2, 64)
}

func HashAndSalt(pwd string) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd+os.Getenv("BCRYPT_SALT")), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd+os.Getenv("BCRYPT_SALT")))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

var charectors = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = charectors[rand.Intn(len(charectors))]
	}
	return string(b)
}
