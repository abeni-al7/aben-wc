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

func (fio FileIO) GetFileSize(path string) int64 {
	fileSize, err := fio.Fs.GetFileSize(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	return fileSize
}

func (fio FileIO) GetLineCount(path string) int {
	lines, err := fio.Fs.GetLineCount(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	return lines
}

func (fio FileIO) GetWordCount(path string) int {
	words, err := fio.Fs.GetWordCount(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	return words
}

func (fio FileIO) GetCharCount(path string) int {
	chars, err := fio.Fs.GetCharCount(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	return chars
}

func (fio FileIO) AcceptInput() {
	byteCount := flag.Bool("c", false, "print byte count")
	lineCount := flag.Bool("l", false, "print line count")
	wordCount := flag.Bool("w", false, "print word count")
	charCount := flag.Bool("m", false, "print character count")

	flag.Parse()
	
	if flag.NArg() != 1 {
		fmt.Println("Usage: abenwc -<arg> <file>")
		os.Exit(1)
	}

	path := flag.Arg(0)
	if *byteCount {
		fileSize := fio.GetFileSize(path)
		fmt.Printf("%d %s\n", fileSize, path)
	} else if *lineCount {
		lines := fio.GetLineCount(path)
		fmt.Printf("%d %s\n", lines, path)
	} else if *wordCount {
		words := fio.GetWordCount(path)
		fmt.Printf("%d %s\n", words, path)
	} else if *charCount {
		chars := fio.GetCharCount(path)
		fmt.Printf("%d %s\n", chars, path)
	} else {
		fileSize := fio.GetFileSize(path)
		lines := fio.GetLineCount(path)
		words := fio.GetWordCount(path)
		fmt.Printf("%d  %d %d %s\n", lines, words, fileSize, path)
	}
}