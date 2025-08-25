package ch10

import (
	"crypto/sha256"
	"io"
	"log"
	"os"
)

func FHash(filepath string) []byte {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	sha := sha256.New()
	_, err = io.Copy(sha, file)
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	return sha.Sum(nil)
}
