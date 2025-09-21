package encryption

import (
	"fmt"
	"os"
)

func CreateNewFile(fileName string, ciphertext []byte) error {
	// WriteFile writes data to the named file, creating it if necessary.
	// If the file does not exist, WriteFile creates it with permissions perm (before umask);
	// otherwise WriteFile truncates it before writing, without changing permissions.
	// Since WriteFile requires multiple system calls to complete, a failure mid-operation can
	// leave the file in a partially written state.
	// https://pkg.go.dev/os#example-WriteFile
	err := os.WriteFile(filePath, data, permission)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	return nil
}
