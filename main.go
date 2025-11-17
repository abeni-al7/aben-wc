package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if (args[0] == "-c") {
		path := args[1]
		fileSize, err := getFileSize(path)
		if err != nil {
			println("There was a problem getting the size of the file")
			return
		}
		fmt.Printf("%d %s\n", fileSize, path)
	}
}

func getFileSize(path string) (int64, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}