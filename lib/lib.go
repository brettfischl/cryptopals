package lib

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
	"os"
)

var LetterFrequency = map[string]float64{
	"a": .08167,
	"b": .01492,
	"c": .02782,
	"d": .04253,
	"e": .12702,
	"f": .02228,
	"g": .02015,
	"h": .06094,
	"i": .06094,
	"j": .00153,
	"k": .00772,
	"l": .04025,
	"m": .02406,
	"n": .06749,
	"o": .07507,
	"p": .01929,
	"q": .00095,
	"r": .05987,
	"s": .06327,
	"t": .09056,
	"u": .02758,
	"v": .00978,
	"w": .02360,
	"x": .00150,
	"y": .01974,
	"z": .00074,
	" ": .13000,
}

func ReadFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scan := bufio.NewScanner(file)
	bytes := make([]byte, 0, 1024)

	for scan.Scan() {
		input := scan.Bytes()
		bytes = append(bytes, input...)
	}
	return bytes
}

func ScoreString(str string) float64 {
	var score float64
	for _, char := range str {
		score += LetterFrequency[string(char)]
	}
	return score
}

func SingleXOR(bytes []byte, single byte) []byte {
	finalBytes := make([]byte, len(bytes))
	for i := range bytes {
		finalBytes[i] = bytes[i] ^ single
	}
	return finalBytes
}

func StringFromSingleByteXOR(bytes []byte, length int) (string, byte) {
	var topScore float64
	var topByte int
	for i := 0; i < length; i++ {
		decodedBytes := SingleXOR(bytes, byte(i))
		score := ScoreString(string(decodedBytes))
		if score > topScore {
			topScore = score
			topByte = i
		}
	}

	return string(SingleXOR(bytes, byte(topByte))), byte(topByte)
}

func FixedXOR(bytes1, bytes2 []byte) ([]byte, error) {
	if len(bytes1) != len(bytes2) {
		return nil, errors.New("Bytes are different lengths")
	}

	finalBytes := make([]byte, len(bytes1))
	for i := range bytes1 {
		finalBytes[i] = bytes1[i] ^ bytes2[i]
	}

	return finalBytes, nil
}

func HexToBase64(str string) string {
	bytes, _ := hex.DecodeString(str)
	base64String := base64.StdEncoding.EncodeToString(bytes)

	return base64String
}

func DecryptRotatingKeyXOR(cipherBytes []byte, key string) string {
	decryptedBytes := make([]byte, len(cipherBytes))

	keyBytes := []byte(key)
	keyIdx := 0
	for i := range cipherBytes {
		decryptedBytes[i] = cipherBytes[i] ^ keyBytes[keyIdx]

		keyIdx++
		if keyIdx == len(keyBytes) {
			keyIdx = 0
		}
	}

	return string(decryptedBytes)
}

func RepeatingKeyXOR(bytes []byte, key string) []byte {
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
