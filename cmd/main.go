package main

import (
	"aes/internal/encryption"
	"aes/internal/input"
	"fmt"
	"os"
)

func main() {
	// No errors required as these are handled in the respective file
	fileName, err := input.ReadInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileContent, err := encryption.ReadContent(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ciphertext, err := encryption.EncryptContent(fileContent)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	errCreate := encryption.CreateNewFile(fileName, ciphertext)
	if errCreate != nil {
		fmt.Printf("failed to create new file: %v", errCreate)
		os.Exit(1)
	}
}
