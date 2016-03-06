package main

import (
  "fmt"
  "flag"
  "os"
  "strconv"
)

func LogVerbose(format string, args ...interface{}) {
  set, _ := strconv.ParseBool(flag.Lookup("verbose").Value.String())
  if set {
    fmt.Fprintf(os.Stderr, format, args...)
  }
}
