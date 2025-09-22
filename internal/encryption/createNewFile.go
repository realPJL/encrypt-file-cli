package encryption

import (
	"fmt"
	"os"
)

func CreateNewFile(fileName string, ciphertext []byte) error {
	// https://pkg.go.dev/os#example-WriteFile
	filePath := fmt.Sprintf("encryptedFiles/ENCRYPTED_%s.enc", fileName)

	errDIR := os.MkdirAll("encryptedFiles", 0755) // permission: owner rwx, group/others rx
	if errDIR != nil {
		return fmt.Errorf("failed to create directory: %w", errDIR)
	}

	data := ciphertext

	err := os.WriteFile(filePath, data, 0644) // permission: owner rw, group/others r
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	return nil
}
