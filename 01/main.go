package main

import (
	"bufio"
	"fmt"
	"os"

	"cryptopals/lib"
)

func main() {
	fmt.Println("Enter a hexidecimal string:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := lib.HexToBase64(scanner.Text())
		fmt.Println(str)
		return
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
