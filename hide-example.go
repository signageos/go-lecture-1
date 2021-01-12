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

	encryptAll(key, exampleNumber)
}

func encryptAll(key []byte, dirPath string) {
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
		if filePath[len(filePath)-4:] == ".enc" {
			continue
		}
		if fileInfo.IsDir() {
			encryptAll(key, filePath)
			continue
		}
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		encryptedContent, err := enc.Encrypt(key, content)
		if err != nil {
			panic(err)
		}
		ioutil.WriteFile(filePath+".enc", encryptedContent, 0644)
	}
}
