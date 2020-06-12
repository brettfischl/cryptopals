package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func repeatingKeyXOR(bytes []byte, key string) []byte {
	var cipherBytes []byte
	keyBytes := []byte(key)
	keyLength := len(keyBytes)
	keyIdx := 0
	for i := range bytes {
		newCipherBytes := bytes[i] ^ keyBytes[keyIdx]
		cipherBytes = append(cipherBytes, newCipherBytes)
		keyIdx += 1
		if keyIdx == keyLength {
			keyIdx = 0
		}
	}

	return cipherBytes
}

func main() {
	var key string
	fmt.Println("Enter your key:")
	fmt.Scanln(&key)

	file, err := os.Open("./05/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	newLineBytes := []byte("\n")
	var cipherBytes []byte
	for scanner.Scan() {
		plainTextBytes := append(scanner.Bytes(), newLineBytes...)
		cipherBytes = repeatingKeyXOR(plainTextBytes, key)
	}

	cipherText := hex.EncodeToString(cipherBytes)

	fmt.Println(string(cipherText))
	return
}
