package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"path/filepath"
	"strings"
	"os"

	"github.com/fatih/color"
)

func SizeStringToBytes() int64 {
	minSizeStr := flag.Lookup("min-size").Value.String()
	minSizeStr = strings.ToLower(minSizeStr)

	var minSizeInt int64
	var d int64
	var s string

	fmt.Sscanf(minSizeStr, "%d%s", &d, &s)

	switch s {
	case "b":
		minSizeInt = d
	case "k":
		minSizeInt = d * 1024
	case "m":
		minSizeInt = d * 1024 * 1024
	case "g":
		minSizeInt = d * 1024 * 1024 * 1024
	default:
		panic("Error reading min-size value")
	}

	return minSizeInt
}

func init() {
	currentDirectory, err := os.Getwd()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to determine the current directory")
		currentDirectory = "/"
	}

	flag.String("dir", currentDirectory , "Directories to scan")
	flag.String("min-size", "1G", "Minimum file size to consider")
	flag.Bool("verbose", false, "Verbose logging to stdout")
}

func main() {
	flag.Parse()
	roots := flag.Lookup("dir").Value.String()
	files.minSize = SizeStringToBytes()

	files.fileList = make([]string, 10)
	files.sizes = make(map[int64][]string, 10)

	for _, r := range strings.Split(roots, " ") {
		red := color.New(color.FgRed).PrintfFunc()
		red("Scanning: %s\n", r)
		err := filepath.Walk(r, visit)

		if err != nil {
			fmt.Printf("filepath.Walk() returned %v\n", err)
		}

	}

	fmt.Println("Done")

	for size, v := range files.sizes {

		if len(v) > 1 {

			green := color.New(color.FgGreen).PrintfFunc()
			tmp := make(map[string][]string, 10)

			for i := 0; i < len(v); i++ {
				hash, err := ComputeMD5(v[i])

				if err != nil {
					fmt.Fprintf(os.Stderr, "Unable to calculate md5 of: %s", v[i])
					continue
				}


				hashHex := hex.EncodeToString(hash)
				tmp[hashHex] = append(tmp[hashHex], v[i])

			}

			for k := range tmp {
				if len(tmp[k]) > 1 {
					green("%dM, %s\n", size/(1024*1024), k)

					for _, files := range tmp[k] {
						green("%s\n", files)
					}

					fmt.Println()

				}
			}
		}
	}
}
