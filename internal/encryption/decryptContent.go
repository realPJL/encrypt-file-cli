package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func DecryptContent(fileContent []byte, keyFileContent []byte) ([]byte, error) {
	// key is being read from user file - subject to change
	key := keyFileContent
	if len(key) != 32 {
		return nil, fmt.Errorf("invalid key length: must be 32 bytes")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesgcm.NonceSize()
	if len(fileContent) < nonceSize {
		return nil, fmt.Errorf("error: Encrypted data is too short (missing nonce)")
	}
	nonce := fileContent[:nonceSize]
	ciphertext := fileContent[nonceSize:]

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("error: Failed to encrypt data: %w", err)
	}

	return plaintext, nil
}
