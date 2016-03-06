package main
import (
  "testing"
  "flag"
  "fmt"
  "math"
  "os"
  "encoding/hex"
)

var testDirectory string

func TestMain(t *testing.T) {
    testDirectory, _ = os.Getwd()
    testDirectory = testDirectory + "/tests/test1"
    flag.Set("dir", testDirectory)
    flag.Set("min-size", "1b")
    main()

    fmt.Println(files.fileList)
}

func TestMD5(t *testing.T) {
  hash_bytes, _ := ComputeMD5(testDirectory + "/testfile.1")
  hash_str := hex.EncodeToString(hash_bytes)
  if hash_str != "d866522038b447a2951dab80ec7f7542" {
    t.Error("Invalid hash returned")
    }
}

func TestSizeStringToBytes(t *testing.T) {
  flag.Set("min-size", "1b")
  bytes := int64(1)
  if (SizeStringToBytes() != bytes) {
    t.Error("1b !=", bytes)
  }

  flag.Set("min-size", "1kb")
  bytes = 1024
  if (SizeStringToBytes() != bytes) {
    t.Error("1kb !=", bytes)
  }

  flag.Set("min-size", "1mb")
  bytes = int64(math.Pow(1024, 2))
  if (SizeStringToBytes() != bytes) {
    t.Error("1mb !=", bytes)
  }

  flag.Set("min-size", "1gb")
  bytes = int64(math.Pow(1024, 3))
  if (SizeStringToBytes() != bytes) {
    t.Error("1GB !=", bytes)
  }

  flag.Set("min-size", "3gb")
  bytes = int64(math.Pow(1024, 3) * 3)
  if (SizeStringToBytes() != bytes) {
    t.Error("1.2GB !=", bytes)
  }

}
