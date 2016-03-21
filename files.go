package main

import "os"

type Files struct {
	fileList []string
	sizes    map[int64][]string
	minSize  int64
}

var files = Files{fileList: []string{}, sizes: map[int64][]string{}}

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
