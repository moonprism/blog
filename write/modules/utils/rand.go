package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func Rand(n int) int {
	return rand.Intn(n)
}

func RandR(min, max int) int {
	return Rand(max-min)+min
}

func RandByte() byte {
	return letterBytes[rand.Intn(len(letterBytes))]
}
