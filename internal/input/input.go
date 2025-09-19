package input

import (
	"aes/internal/encryption"
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInput() {
	fmt.Println("Please enter the file name you wish to encrypt: ")

	fileName := bufio.NewScanner(os.Stdin)
	fileName.Scan()
	err := fileName.Err()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Checking if '%s' is there...\n", fileName.Text())

	_, err = os.Stat(fileName.Text())

	if os.IsNotExist(err) {
		fmt.Println("This file does not exist in this directory.")
	} else {
		fmt.Println("SUCCESS \nFile found!")
		encryption.ReadContent(fileName.Text())
	}
}
