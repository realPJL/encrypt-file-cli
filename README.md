# AES-Encrypt-CLI
AES-Encrypt is a CLI-tool to encrypt any file you want. Simply start the program, enter the file name, your key and encrypt it using **AES-256**.  

This tool uses the **GCM (Galois/Counter Mode)** because it allows messages longer than the block length of the block cipher to be encrypted.  

A GUI version is likely.

‼️ Trust, but verify ‼️

THIS IS FOR EDUCATIONAL PURPOSES. DON'T RELY ON IT!!! (it works though)

## Installation

Clone the repository:
```bash
  git clone https://github.com/realPJL/encrypt-file-cli.git
  cd encrypt-file-cli
```

Go to the ```cmd``` folder and build the app (replace ```theApplicationName``` with your desired name):
```bash
  cd cmd
  go build -o theApplicationName
```

Move the new application to the folder containing your ```secret``` and ```key``` file. Keep in mind that the ```key``` needs to be 32 bytes big.

---
### The application, the key and the desired file need to be in the same folder when running the program.

## Usage

On Windows:
```bash
  theApplicationName.exe
```

On macOS, Linux
```bash
  ./theApplicationName
```

After starting the program choose your desired mode ```E``` (Encryption) or ```D``` (Decryption).
Then, please type in the name of the file you want to en-/decrypt:
```bash
  secret.md
```
The file will be read. After it has been read enter your key file:
```bash
  key.txt
````
Your encrypted file will be stored inside a new folder ```encryptedFiles``` and have the name of your original file and the prefix ```ENCRYPTED_``` as well as a new file ending ```.enc```.

## Contributing

Feel free to contribute your desired functionality. Just fork the repo and start working. I'll probably abon this CLI version for the GUI version (it's just better UX for normal people). You can find the GUI version here: TBD

## License

MIT License (see LICENSE file for more information)