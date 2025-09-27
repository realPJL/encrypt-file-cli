package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("----------WIP-AES-Tool----------")
	fmt.Println("Do you want to Encrypt or Decrypt a file? (E/D): ")

	mode := bufio.NewScanner(os.Stdin)
	mode.Scan()
	err := mode.Err()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch strings.ToUpper(mode.Text()) {
	case "E":
		runEncryption()
	case "D":
		runDecryption()
	case "":
		runEncryption() // Should be default
	default:
		runEncryption()
	}
}
