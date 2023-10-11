package random

import (
	"math/rand"
	"time"
)

const englishAlphabet = "abcdefjhigklmnopqrstuvwxyz"

func NewRandomString(length int) string {
	randomString := ""
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		randomString += englishAlphabet[rand.Intn(len(englishAlphabet))]
	}
	return randomString
}
