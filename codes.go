package main

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHJKLMNOPQRSTUVWXYZ023456789"
const codeLength = 5
const numCodes = 100

func GenerateRandomCode(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(code)
}

func GenerateUniqueCodes(numCodes int, codeLength int) []string {
	uniqueCodes := make(map[string]bool)
	codes := []string{}

	for len(uniqueCodes) < numCodes {
		code := GenerateRandomCode(codeLength)
		uniqueCodes[code] = true
	}

	for code := range uniqueCodes {
		codes = append(codes, code)
	}

	return codes
}

func main() {
	uniqueCodes := GenerateUniqueCodes(numCodes, codeLength)

	fmt.Println("Generated Unique Codes:")
	for _, code := range uniqueCodes {
		fmt.Println(code)
	}
}
