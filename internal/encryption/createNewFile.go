package encryption

import (
	"fmt"
	"os"
	"strings"
)

func CreateNewFile(fileName string, ciphertext []byte, mode string) error {
	switch mode {
	case "E":
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
	case "D":
		s := fileName
		fileNameClean := strings.Trim(s, ".enc")

		// https://pkg.go.dev/os#example-WriteFile
		filePath := fmt.Sprintf("decryptedFiles/DECRYPTED_%s", fileNameClean)

		errDIR := os.MkdirAll("decryptedFiles", 0755) // permission: owner rwx, group/others rx
		if errDIR != nil {
			return fmt.Errorf("failed to create directory: %w", errDIR)
		}

		data := ciphertext

		err := os.WriteFile(filePath, data, 0644) // permission: owner rw, group/others r
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("error: Something went wrong while creating the new file. Please try again")
	}
}
