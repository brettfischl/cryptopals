package main

import (
	"cryptopals/lib"
	"encoding/base64"
	"fmt"
)

const MIN_KEY_SIZE = 2
const MAX_KEY_SIZE = 40

func HammingDistance(bytes1, bytes2 []byte) int {
	differingBytes := 0
	for bit := range bytes1 {
		diff := bytes1[bit] ^ bytes2[bit]
		for diff != 0 {
			if diff&1 == 1 {
				differingBytes++
			}
			diff = diff >> 1
		}
	}
	return differingBytes
}

func findKeySize(bytes []byte) int {
	minDist := 1000.00
	bestKeySize := 0
	keySizes := make(map[int]float64)
	for keySize := MIN_KEY_SIZE; keySize <= MAX_KEY_SIZE; keySize++ {
		cumulativeDist := 0.0
		blocks := 0
		for blockIdx := 0; blockIdx+(keySize*2) < len(bytes); blockIdx += keySize {
			block1 := bytes[blockIdx : blockIdx+keySize]
			block2 := bytes[blockIdx+keySize : blockIdx+(keySize*2)]
			hammingDistance := HammingDistance(block1, block2)
			cumulativeDist += float64(hammingDistance) / float64(keySize)
			blocks++
		}

		normalizedDist := float64(cumulativeDist) / float64(blocks)
		keySizes[keySize] = normalizedDist
		if normalizedDist < minDist {
			minDist = normalizedDist
			bestKeySize = keySize
		}
	}
	return bestKeySize
}

func generateBlocks(cipherText []byte, keySize int) map[int][]byte {
	blocks := make(map[int][]byte)
	for i, byte := range cipherText {
		pos := i % keySize
		blocks[pos] = append(blocks[pos], byte)
	}
	return blocks
}

func breakRepeatingKeyXOR(cipherText []byte, keySize int) []byte {
	blocks := generateBlocks(cipherText, keySize)
	finalKey := make([]byte, keySize)

	for i, val := range blocks {
		_, key := lib.StringFromSingleByteXOR(val, 256)
		finalKey[i] = key
	}
	return finalKey
}

func main() {
	encodedBytes := lib.ReadFile("./06/input.txt")
	bytes, _ := base64.StdEncoding.DecodeString(string(encodedBytes))

	keySize := findKeySize(bytes)

	key := breakRepeatingKeyXOR(bytes, keySize)

	final := lib.RepeatingKeyXOR(bytes, string(key))
	fmt.Println("PLAINTEXT:\n", string(final))
	fmt.Println("KEY:\n", string(key))
	fmt.Println("KEYSIZE:\n", keySize)
	return
}
