package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	byteCount := flag.Bool("c", false, "print byte count")
	flag.Parse()
	
	if !*byteCount || flag.NArg() != 1 {
		fmt.Println("Usage: abenwc -c <file>")
		os.Exit(1)
	}
	
	path := flag.Arg(0)
	fileSize, err := getFileSize(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d %s\n", fileSize, path)
}

func getFileSize(path string) (int64, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}