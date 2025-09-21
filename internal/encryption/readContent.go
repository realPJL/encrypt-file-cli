package encryption

import (
	"fmt"
	"log"
	"os"
)

func ReadContent(fileName string) {
	fmt.Println("The file is being read... Please wait...")

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatalf("failed to get file info: %v", err)
	}

	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("failed to read file content: %v", err)
	}

	fmt.Printf("File Size: %v bytes\n", fileInfo.Size())
	// NOTE: output is a []byte slice. Will not be readable if anything but text
	// Uncomment for debug purposes
	//fmt.Printf("--------------THE-CONTENT-------------------------\n%s\n", fileContent)
	//fmt.Println("--------------THE-END-------------------------")

	EncryptContent(fileContent)
}
