package main

import (
	"aes/internal/encryption"
	"aes/internal/input"
	"log"
)

func main() {
	fileName, errInput := input.ReadInput()
	if errInput != nil {
		log.Fatalf("404")
	}

	fileContent, errContent := encryption.ReadContent(fileName)
	if errContent != nil {
		log.Fatalf("Error while reading file content")
	}

	ciphertext, errEncrypt := encryption.EncryptContent(fileContent)
	if errEncrypt != nil {
		log.Fatalf("Error while encrypting the content")
	}

	encryption.CreateNewFile(fileName, ciphertext)
}
