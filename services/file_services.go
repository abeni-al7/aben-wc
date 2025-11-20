package services

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

type FileService struct{}


func (fs FileService) GetFileSize(data []byte) (int) {
	return len(data)
}

func (fs FileService) GetLineCount(data []byte) (int) {
	return bytes.Count(data, []byte{'\n'})
}

func (fs FileService) GetWordCount(data []byte) (int) {
	return len(strings.Fields(string(data)))
}

func (fs FileService) GetCharCount(data []byte) (int) {
	return utf8.RuneCount(data)
}
