package services

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

type FileService struct{}


func (fs FileService) GetFileSize(data []byte) (int, error) {
	return len(data), nil
}

func (fs FileService) GetLineCount(data []byte) (int, error) {
	return bytes.Count(data, []byte{'\n'}), nil
}

func (fs FileService) GetWordCount(data []byte) (int, error) {
	return len(strings.Fields(string(data))), nil
}

func (fs FileService) GetCharCount(data []byte) (int, error) {
	return utf8.RuneCount(data), nil
}
