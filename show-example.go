package main

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/signageos/go-lecture-1/enc"
)

func main() {
	exampleNumber := os.Args[1]
	password := os.Args[2]

	key, _, err := enc.DeriveKey([]byte(password), []byte("sůl"))
	if err != nil {
		panic(err)
	}

	decryptAll(key, exampleNumber)
}

func decryptAll(key []byte, dirPath string) {
	dir, err := os.Open(dirPath)
	if err != nil {
		panic(err)
	}
	files, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}
	for _, fileInfo := range files {
		filePath := path.Join(dirPath, fileInfo.Name())
		if filePath[len(filePath)-4:] != ".enc" {
			continue
		}
		if fileInfo.IsDir() {
			decryptAll(key, filePath)
			continue
		}
		encryptedContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		content, err := enc.Decrypt(key, encryptedContent)
		if err != nil {
			panic(err)
		}
		ioutil.WriteFile(filePath[:len(filePath)-4], content, 0644)
	}

}
