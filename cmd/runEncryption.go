package main

import (
	"aes/internal/encryption"
	"aes/internal/input"
	"fmt"
	"os"
)

func runEncryption() {
	fmt.Println("----------AES-ENCRYPTION----------")
	fileName, err := input.ReadInput("E")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileContent, err := encryption.ReadContent(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	keyFileName, err := input.ReadInput("K")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	keyFileContent, err := encryption.ReadContent(keyFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	encryptedData, err := encryption.EncryptContent(fileContent, keyFileContent)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	errCreate := encryption.CreateNewFile(fileName, encryptedData, "E")
	if errCreate != nil {
		fmt.Printf("failed to create new file: %v", errCreate)
		os.Exit(1)
	}
}
