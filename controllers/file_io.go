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

	counts, err := fio.Fs.CalculateCounts(r)
	if err != nil {
		fio.handleError(err)
	}

	fio.printOutput(counts, path, flags)
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

func (fio FileIO) handleError(err error) {
	if err.Error() == "usage: abenwc -<arg> <filepath>" {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	os.Exit(1)
}

func (fio FileIO) printOutput(counts services.Counts, path string, flags WcFlags) {
	if flags.ByteCount {
		fmt.Printf("%d %s\n", counts.Bytes, path)
	} else if flags.LineCount {
		fmt.Printf("%d %s\n", counts.Lines, path)
	} else if flags.WordCount {
		fmt.Printf("%d %s\n", counts.Words, path)
	} else if flags.CharCount {
		fmt.Printf("%d %s\n", counts.Chars, path)
	} else {
		fmt.Printf("%d  %d %d %s\n", counts.Lines, counts.Words, counts.Bytes, path)
	}
}