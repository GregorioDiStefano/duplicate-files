package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"errors"
	"fmt"
	"hash"
	"io"
	"os"
)

func DoHashCalculation(hasher hash.Hash, filePath string) ([]byte, error) {

	var result []byte
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		fmt.Println("Failed opening:", filePath)
		return result, err
	}

	if err != nil {
		fmt.Println("md5 not computed - file failed to open or is not regular:", filePath)
		return result, errors.New("Error")
	}

	if _, err := io.Copy(hasher, file); err != nil {
		return result, err
	}

	return hasher.Sum(result), nil
}

func ComputeHash(hashType, filePath string) ([]byte, error) {
	switch hashType {
	case "md5":
		return DoHashCalculation(md5.New(), filePath)
	case "sha1":
		return DoHashCalculation(sha1.New(), filePath)
	case "sha256":
		return DoHashCalculation(sha512.New(), filePath)
	case "sha512":
		return DoHashCalculation(sha512.New(), filePath)
	default:
		return DoHashCalculation(md5.New(), filePath)
	}
}
