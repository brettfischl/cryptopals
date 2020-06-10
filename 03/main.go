package main

import (
	"cryptopals/lib"
	"encoding/hex"
	"fmt"
)

// https://en.wikipedia.org/wiki/Letter_frequency

func scoreString(str string) float64 {
	var score float64
	fmt.Println(str)
	for _, char := range str {
		score += lib.LetterFrequency[string(char)]
	}
	return score
}

func singleXOR(bytes []byte, single byte) []byte {
	finalBytes := make([]byte, len(bytes))
	for i := range bytes {
		finalBytes[i] = bytes[i] ^ single
	}
	return finalBytes
}

func main() {
	var input string
	fmt.Println("Enter your first string:")
	fmt.Scanln(&input)
	bytes, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}

	finalStr := lib.StringFromSingleByteXOR(bytes, 256)
	fmt.Println(finalStr)
	return
}
