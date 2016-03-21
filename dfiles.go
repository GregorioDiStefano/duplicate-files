package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func SizeStringToBytes() int64 {
	minSizeStr := flag.Lookup("min-size").Value.String()
	minSizeStr = strings.ToLower(minSizeStr)

	var minSizeInt, d int64
	var s string

	fmt.Sscanf(minSizeStr, "%d%s", &d, &s)

	switch s {
	case "b":
		minSizeInt = d
	case "k", "kb":
		minSizeInt = d * 1024
	case "m", "mb":
		minSizeInt = d * 1024 * 1024
	case "g", "gb":
		minSizeInt = d * 1024 * 1024 * 1024
	default:
		panic("Error reading min-size value: " + s)
	}

	LogDebug("minSize for file set to: %d\n", minSizeInt)
	return minSizeInt
}

func init() {
	currentDirectory, err := os.Getwd()

	if err != nil {
		LogVerbose("red", "Unable to determine the current directory")
		os.Exit(1)
	}

	flag.String("dir", currentDirectory, "Directories to scan")
	flag.String("min-size", "1G", "Minimum file size to consider")
	flag.String("hash-type", "md5", "Hash type: md5, sha1")
	flag.String("save", "", "Save file information to file")
	flag.Bool("verbose", false, "Verbose logging to stdout")
	flag.Bool("debug", false, "Debug logging to stdout")
}

func printDuplicateFiles(duplicates map[string][]string, size int64) {
	for k := range duplicates {
		if len(duplicates[k]) > 1 {
			DefaultPrint("green", "%dM, %s\n", size/(1024*1024), k)

			for _, files := range duplicates[k] {
				DefaultPrint("green", "%s\n", files)
			}
			DefaultPrint(nil, "\n")
		}
	}
}

func SaveToJSON(hashType string) {
	sizes_to_files := make(map[string][]string, 10)
	for size, files := range files.sizes {
		files_array := []string{"", "", "", ""}

		copy(files_array[:], files[0:len(files)])
		size_as_string := strconv.FormatInt(size, 10)
		for _, f := range files {
			sizes_to_files[size_as_string] = append(sizes_to_files[size_as_string],
				f)
		}
	}
	m, _ := json.Marshal(sizes_to_files)
	fmt.Println(m, sizes_to_files)
}

func PrintDuplicateFiles(hashType string) {
	for size, v := range files.sizes {

		if len(v) > 1 {
			tmp := make(map[string][]string, 10)

			for i := 0; i < len(v); i++ {
				hash, err := ComputeHash(hashType, v[i])

				if err != nil {
					continue
				}

				hashHex := hex.EncodeToString(hash)
				tmp[hashHex] = append(tmp[hashHex], v[i])
			}
			printDuplicateFiles(tmp, size)
		}
	}
}

func main() {
	flag.Parse()
	roots := flag.Lookup("dir").Value.String()
	hashType := flag.Lookup("hash-type").Value.String()

	files.minSize = SizeStringToBytes()
	files.fileList = make([]string, 10)
	files.sizes = make(map[int64][]string, 10)

	saveFile := flag.Lookup("save").Value.String()

	for _, r := range strings.Split(roots, " ") {
		DefaultPrint("green", "Scanning: %s\n", r)
		err := filepath.Walk(r, visit)

		if err != nil {
			fmt.Printf("filepath.Walk() returned %v\n", err)
		}
	}
	if len(saveFile) == 0 {
		DefaultPrint(nil, "Done\n")
		PrintDuplicateFiles(hashType)
	} else {
		SaveToJSON(hashType)
	}
}
