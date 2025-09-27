package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInput(mode string) (string, error) {

	switch mode {
	case "E":
		fmt.Println("Please enter the file name you wish to encrypt: ")
	case "D":
		fmt.Println("Please enter the file name you wish to decrypt: ")
	default:
		fmt.Println("Please enter the file name you wish to encrypt: ")
	}

	fileName := bufio.NewScanner(os.Stdin)
	fileName.Scan()
	err := fileName.Err()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Checking if '%s' is there...\n", fileName.Text())

	_, err = os.Stat(fileName.Text())

	if os.IsNotExist(err) {
		return "", fmt.Errorf("this file does not exist in this directory")
	} else {
		fmt.Println("SUCCESS \nFile found!")
	}

	return fileName.Text(), nil
}
