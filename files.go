package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

type Files struct {
	fileList []string
	sizes    map[int64][]string
	minSize  int64
}

var files = Files{fileList: []string{}, sizes: map[int64][]string{}}

func ComputeMD5(filePath string) ([]byte, error) {
	var result []byte
	hash := md5.New()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed opening:", filePath)
		return result, err
	}

	defer file.Close()

	if _, err := io.Copy(hash, file); err != nil {
		return result, err
	}

	return hash.Sum(result), nil
}

func visit(path string, f os.FileInfo, err error) error {
	if err != nil || f.Size() == 0 {
		return nil
	}

	if f.Size() >= files.minSize {
		fmt.Printf("Visited: %s, %d\n", path, f.Size())

		if files.sizes == nil {
			files.sizes = make(map[int64][]string, 10)
			fmt.Println("Here")
		}
		files.sizes[f.Size()] = append(files.sizes[f.Size()], path)
		files.fileList = append(files.fileList, path)
	}

	return nil
}
