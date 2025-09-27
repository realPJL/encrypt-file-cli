package main

import (
	"aes/internal/encryption"
	"aes/internal/input"
	"fmt"
	"os"
)

func runDecryption() {
	fmt.Println("----------AES-DECRYPTION----------")
	fileName, err := input.ReadInput("D")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileContent, err := encryption.ReadContent(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	plaintext, err := encryption.DecryptContent(fileContent)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	errCreate := encryption.CreateNewFile(fileName, plaintext, "D")
	if errCreate != nil {
		fmt.Printf("failed to create new file: %v", errCreate)
		os.Exit(1)
	}
}
