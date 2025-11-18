package controllers

import (
	"flag"
	"fmt"
	"os"

	"github.com/abeni-al7/aben-wc/services"
)

type FileIO struct {
	Fs services.FileService
}

func (fio FileIO) AcceptInput() {
	byteCount := flag.Bool("c", false, "print byte count")
	flag.Parse()
	
	if !*byteCount || flag.NArg() != 1 {
		fmt.Println("Usage: abenwc -c <file>")
		os.Exit(1)
	}
	
	path := flag.Arg(0)
	fileSize, err := fio.Fs.GetFileSize(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d %s\n", fileSize, path)
}