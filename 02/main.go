package main

import (
	"cryptopals/lib"
	"encoding/hex"
	"fmt"
)

func main() {
	var string1 string
	var string2 string
	fmt.Println("Enter your first string:")
	fmt.Scanln(&string1)
	fmt.Println("Enter your second string:")
	fmt.Scanln(&string2)

	bytes1, err := hex.DecodeString(string1)
	bytes2, err := hex.DecodeString(string2)

	if err != nil {
		panic(err)
	}

	xorBytes, err := lib.FixedXOR(bytes1, bytes2)
	if err != nil {
		panic(err)
	}

	str := hex.EncodeToString(xorBytes)
	fmt.Println(str)
	return
}
