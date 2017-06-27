package main

import (
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

var (
	info = color.New(color.FgYellow).SprintfFunc()
)

const (
	perfc    = "C:\\Windows\\perfc"
	perfcDLL = "C:\\Windows\\perfc.dll"
	perfcDAT = "C:\\Windows\\perfc.dat"

	READONLY = os.FileMode(os.O_RDONLY)
)

func main() {
	if _, err := os.Stat(perfc); os.IsNotExist(err) {
		ioutil.WriteFile(perfc, []byte{}, READONLY)
	} else {
		info("File already exists: %s\nNo action needed", perfc)
	}
	if _, err := os.Stat(perfcDLL); os.IsNotExist(err) {
		ioutil.WriteFile(perfcDLL, []byte{}, READONLY)
	} else {
		info("File already exists: %s\nNo action needed", perfcDLL)
	}
	if _, err := os.Stat(perfcDAT); os.IsNotExist(err) {
		ioutil.WriteFile(perfcDAT, []byte{}, READONLY)
	} else {
		info("File already exists: %s\nNo action needed", perfcDAT)
	}
}
