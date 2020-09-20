package main

import (
	"cryptopals/lib"

	"encoding/base64"
	"fmt"
)

func main() {
	encodedString := lib.ReadFile("./07/input.txt")
	bytes, _ := base64.StdEncoding.DecodeString(string(encodedString))

	key := []byte("YELLOW SUBMARINE")
	plainText := lib.DecryptAES128ECB(bytes, key)

	fmt.Println("PLAINTEXT:\n", string(plainText))
	fmt.Println("KEY:\n", string(key))
	return
}
