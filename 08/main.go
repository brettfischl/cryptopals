package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func DetectAES128ECB(cipherText []byte, keyLength int) bool {
	for i := 0; i < len(cipherText); i += 1 {
		for j := i + 1; j < len(cipherText); j += 1 {
			blockAIdxStart := i
			blockAIdxEnd := i + keyLength

			blockBIdxStart := j
			blockBIdxEnd := j + keyLength

			blockA := cipherText[blockAIdxStart:blockAIdxEnd]
			blockB := cipherText[blockBIdxStart:blockBIdxEnd]
			if bytes.Equal(blockA, blockB) {
				return true
			}
		}
	}
	return false
}

func main() {
	file, err := os.Open("./08/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cipherText, err := hex.DecodeString(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		detected := DetectAES128ECB(cipherText, 16)
		if detected {
			fmt.Println(hex.EncodeToString(cipherText))
			break
		}

		// blockLength := 16

		// blocks := [][]byte{}
		// for i := 0; i < len(cipherText); i += blockLength {
		// 	blocks := append(blocks, cipherText[i:i+blockLength])
		// }

		// for i, blockA := range blocks {
		// 	for j, blockB := range blocks {
		// 		if i != j && bytes.Equal(blockA, blockB) {
		// 			fmt.Println(string(cipherText))
		// 		}
		// 	}
		// }
	}

	fmt.Println("done")
}
