package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No input file")
		os.Exit(1)
	}
	inputFileName := os.Args[1]

	pFile, err := os.Open(inputFileName)

	if err != nil {
		fmt.Println("Error occurred while reading the file")
		os.Exit(2)
	}

	io.Copy(os.Stdout, pFile)
}
