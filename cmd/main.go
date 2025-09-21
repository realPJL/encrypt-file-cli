package main

import (
	"aes/internal/encryption"
	"aes/internal/input"
)

func main() {
	// No errors required as these are handled in the respective file
	fileName, _ := input.ReadInput()

	fileContent, _ := encryption.ReadContent(fileName)

	ciphertext, _ := encryption.EncryptContent(fileContent)

	encryption.CreateNewFile(fileName, ciphertext)
}
