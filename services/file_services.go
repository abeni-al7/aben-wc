package services

import (
	"bufio"
	"io"
	"unicode"
)

type FileService struct{}

type Counts struct {
	Bytes int64
	Lines int
	Words int
	Chars int
}

func (fs FileService) CalculateCounts(r io.Reader) (Counts, error) {
	counts := Counts{}
	reader := bufio.NewReader(r)

	inWord := false

	for {
		r, size, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return counts, err
		}

		counts.Bytes += int64(size)
		counts.Chars++

		if r == '\n' {
			counts.Lines++
		}

		isSpace := unicode.IsSpace(r)
		if inWord && isSpace {
			inWord = false
		} else if !inWord && !isSpace {
			inWord = true
			counts.Words++
		}
	}
	return counts, nil
}
