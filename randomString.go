package main

import (
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandomString(length int, p1 int, p2 int) string {
	if p1+p2 > length {
		panic("p1 + p2 cannot be greater than length")
	}
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	result := make([]byte, length)

	for i := 0; i < p1; i++ {
		result[i] = chars[seededRand.Intn(len(chars))]
	}
	for i := p1; i < p1+p2; i++ {
		result[i] = digits[seededRand.Intn(len(digits))]
	}
	for i := p1 + p2; i < length; i++ {
		charset := chars + digits
		result[i] = charset[seededRand.Intn(len(charset))]
	}

	seededRand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return string(result)
}
