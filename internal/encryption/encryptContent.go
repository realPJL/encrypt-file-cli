package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func EncryptContent(fileContent []byte, keyFileContent []byte) ([]byte, error) {
	// key is being read from user file - subject to change
	key := keyFileContent
	if len(key) != 32 {
		return nil, fmt.Errorf("invalid key length: must be 32 bytes")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	encryptedData := aesgcm.Seal(nonce, nonce, fileContent, nil)

	fmt.Println("Done! Ciphertext generated. Please wait while the encrypted file is being created.")

	return encryptedData, nil
}
