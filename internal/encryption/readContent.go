package encryption

import (
	"fmt"
	"os"
)

func ReadContent(fileName string) ([]byte, error) {
	fmt.Println("The file is being read... Please wait...")

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %v", err)
	}

	fmt.Printf("File Size: %v bytes\n", fileInfo.Size())
	// NOTE: output is a []byte slice. Will not be readable if anything but text
	// Uncomment for debug purposes
	//fmt.Printf("--------------THE-CONTENT-------------------------\n%s\n", fileContent)
	//fmt.Println("--------------THE-END-------------------------")

	return fileContent, nil
}
