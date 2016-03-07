package main

import (
	"crypto/md5"
	"errors"
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
	defer file.Close()

	if err != nil {
		fmt.Println("Failed opening:", filePath)
		return result, err
	}

	if err != nil {
		fmt.Println("md5 not computed - file failed to open or is not regular:", filePath)
		return result, errors.New("Error")
	}

	if _, err := io.Copy(hash, file); err != nil {
		return result, err
	}

	return hash.Sum(result), nil
}

func visit(path string, f os.FileInfo, err error) error {
	if err != nil || f.Size() == 0 || f.IsDir() {
		return nil
	}

	fileMode := f.Mode().IsRegular()

	if f.Size() >= files.minSize && fileMode {
		LogVerbose("green", "Visited: %s, %d\n", path, f.Size())

		files.sizes[f.Size()] = append(files.sizes[f.Size()], path)
		files.fileList = append(files.fileList, path)
	}

	return nil
}
