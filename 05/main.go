package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	var key string
	fmt.Println("Enter your key:")
	fmt.Scanln(&key)
	keyBytes := []byte(key)
	keyLength := len(keyBytes)

	file, err := os.Open("./05/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	newLineBytes := []byte("\n")
	keyIdx := 0
	var cipherBytes []byte
	for scanner.Scan() {
		plainTextBytes := append(scanner.Bytes(), newLineBytes...)

		for i := range plainTextBytes {
			newCipherBytes := plainTextBytes[i] ^ keyBytes[keyIdx]
			cipherBytes = append(cipherBytes, newCipherBytes)
			keyIdx += 1
			if keyIdx == keyLength {
				keyIdx = 0
			}
		}
	}

	cipherText := hex.EncodeToString(cipherBytes)

	fmt.Println(string(cipherText))
	return
}
