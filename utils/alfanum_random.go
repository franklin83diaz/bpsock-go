package utils

import (
	"math/rand"
	"time"
)

// AlfanumRandom generates a random alphanumeric string of length n
func AlfanumRandom(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	source := rand.NewSource(int64(time.Now().UnixNano()))
	rng := rand.New(source)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rng.Intn(len(letters))]
	}
	return string(b)
}
