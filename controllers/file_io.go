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

type WcFlags struct {
	ByteCount bool
	LineCount bool
	WordCount bool
	CharCount bool
}

func (fio FileIO) AcceptInput() {
	flags := fio.parseFlags()

	r, path, err := fio.getInputSource()
	if err != nil {
		fio.handleError(err)
	}
	defer r.Close()

	data, err := fio.readData(r)
	if err != nil {
		fio.handleError(err)
	}

	fio.printOutput(data, path, flags)
}

func (fio FileIO) parseFlags() WcFlags {
	byteCount := flag.Bool("c", false, "print byte count")
	lineCount := flag.Bool("l", false, "print line count")
	wordCount := flag.Bool("w", false, "print word count")
	charCount := flag.Bool("m", false, "print character count")

	flag.Parse()

	return WcFlags{
		ByteCount: *byteCount,
		LineCount: *lineCount,
		WordCount: *wordCount,
		CharCount: *charCount,
	}
}

func (fio FileIO) getInputSource() (*os.File, string, error) {
	stat, err := os.Stdin.Stat()
	if err == nil && (stat.Mode()&os.ModeCharDevice) == 0 {
		return os.Stdin, "", nil
	}

	if flag.NArg() == 1 {
		path := flag.Arg(0)
		f, err := os.Open(path)
		return f, path, err
	}

	return nil, "", fmt.Errorf("usage: abenwc -<arg> <filepath>")
}

func (fio FileIO) readData(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}

func (fio FileIO) handleError(err error) {
	if err.Error() == "usage: abenwc -<arg> <filepath>" {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	os.Exit(1)
}

func (fio FileIO) printOutput(data []byte, path string, flags WcFlags) {
	if flags.ByteCount {
		fileSize := fio.Fs.GetFileSize(data)
		fmt.Printf("%d %s\n", fileSize, path)
	} else if flags.LineCount {
		lines := fio.Fs.GetLineCount(data)
		fmt.Printf("%d %s\n", lines, path)
	} else if flags.WordCount {
		words := fio.Fs.GetWordCount(data)
		fmt.Printf("%d %s\n", words, path)
	} else if flags.CharCount {
		chars := fio.Fs.GetCharCount(data)
		fmt.Printf("%d %s\n", chars, path)
	} else {
		fileSize := fio.Fs.GetFileSize(data)
		lines := fio.Fs.GetLineCount(data)
		words := fio.Fs.GetWordCount(data)
		fmt.Printf("%d  %d %d %s\n", lines, words, fileSize, path)
	}
}