package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

var isTesting = false
var stdOutString []string

func unsetColor() {
	color.Unset()
}

func setColor(strColor interface{}) {
	if strColor == nil {
		color.Unset()
		return
	}

	switch strColor {
	case "red":
		color.Set(color.FgRed)
	case "green":
		color.Set(color.FgGreen)
	default:
		color.Unset()
	}
}

func DefaultPrint(strColor interface{}, format string, args ...interface{}) {
	if isTesting {
		stdOutString = append(stdOutString, fmt.Sprintf(format, args...))
	} else {
		defer unsetColor()
		setColor(strColor)
		fmt.Fprintf(os.Stdout, format, args...)
	}
}

func LogVerbose(strColor interface{}, format string, args ...interface{}) {
	defer unsetColor()
	setColor(strColor)
	set, _ := strconv.ParseBool(flag.Lookup("verbose").Value.String())
	if set {
		fmt.Fprintf(os.Stderr, format, args...)
	}
}

func LogDebug(format string, args ...interface{}) {
	set, _ := strconv.ParseBool(flag.Lookup("debug").Value.String())
	if set {
		t := time.Now()
		date_time := fmt.Sprintf("%02d:%02d:%02d: ", t.Hour(), t.Minute(), t.Second())
		fmt.Fprintf(os.Stderr, date_time+format, args...)
	}
}
