package encryption

import (
	"fmt"
	"log"
	"os"
)

func ReadContent(fileName string) {
	fmt.Println("The file is being read... Please wait...")

	fileSize, err1 := os.Stat(fileName)

	fileContent, err2 := os.ReadFile(fileName)

	if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
	}

	fmt.Printf("File Size: %v bytes\n", fileSize.Size())
	// NOTE: output is a []byte slice. Will not be readable if anything but text
	fmt.Printf("--------------THE-CONTENT-------------------------\n%s\n", fileContent)
	fmt.Println("--------------THE-END-------------------------")
}
