package random

import (
	"math/rand"
)

const englishAlphabet = "abcdefjhigklmnopqrstuvwxyz"

func NewRandomString(length int) string {
	randomString := ""
	for i := 0; i < length; i++ {
		randomString += string(englishAlphabet[rand.Intn(len(englishAlphabet))])
	}
	return randomString
}
