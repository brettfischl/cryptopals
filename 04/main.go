package main

import (
	"bufio"
	"cryptopals/lib"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./04/input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	linesWithScores := make(map[string]float64)
	for scanner.Scan() {
		bytes, err := hex.DecodeString(scanner.Text())
		if err != nil {
			panic(err)
		}

		finalStr := lib.StringFromSingleByteXOR(bytes, 256)
		linesWithScores[finalStr] = lib.ScoreString(finalStr)
	}

	var topScore float64
	var topLine string
	for k, v := range linesWithScores {
		if v > topScore {
			topScore = v
			topLine = k
		}
	}

	fmt.Println(topLine)
	return
}
