package controllers

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/abeni-al7/aben-wc/services"
)

type FileIO struct {
	Fs services.FileService
}

func (fio FileIO) GetFileSize(data []byte) int {
	fileSize, err := fio.Fs.GetFileSize(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	return fileSize
}

func (fio FileIO) GetLineCount(data []byte) int {
	lines, err := fio.Fs.GetLineCount(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	return lines
}

func (fio FileIO) GetWordCount(data []byte) int {
	words, err := fio.Fs.GetWordCount(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	return words
}

func (fio FileIO) GetCharCount(data []byte) int {
	chars, err := fio.Fs.GetCharCount(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	return chars
}

func (fio FileIO) AcceptInput() {
	var r *os.File 
	var path string

	byteCount := flag.Bool("c", false, "print byte count")
	lineCount := flag.Bool("l", false, "print line count")
	wordCount := flag.Bool("w", false, "print word count")
	charCount := flag.Bool("m", false, "print character count")
	
	flag.Parse()

	stat, err := os.Stdin.Stat()
	if (err == nil) && ((stat.Mode() & os.ModeCharDevice) == 0) {
		r = os.Stdin
	} else if flag.NArg() == 1 {
		path = flag.Arg(0)
		fmt.Println(path)
		r, err = os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		defer r.Close()
	} else {
		fmt.Fprintf(os.Stderr, "Usage: abenwc -<arg> <filepath>\n")
		os.Exit(1)
	}

	data, err := io.ReadAll(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if *byteCount {
		fileSize := fio.GetFileSize(data)
		fmt.Printf("%d %s\n", fileSize, path)
	} else if *lineCount {
		lines := fio.GetLineCount(data)
		fmt.Printf("%d %s\n", lines, path)
	} else if *wordCount {
		words := fio.GetWordCount(data)
		fmt.Printf("%d %s\n", words, path)
	} else if *charCount {
		chars := fio.GetCharCount(data)
		fmt.Printf("%d %s\n", chars, path)
	} else {
		fileSize := fio.GetFileSize(data)
		lines := fio.GetLineCount(data)
		words := fio.GetWordCount(data)
		fmt.Printf("%d  %d %d %s\n", lines, words, fileSize, path)
	}
}