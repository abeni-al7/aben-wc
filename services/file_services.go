package services

import (
	"bufio"
	"fmt"
	"os"
)

type FileService struct {}

func (fs FileService) GetFileSize(path string) (int64, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	if !fileInfo.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", path)
	}
	return fileInfo.Size(), nil
}

func (fs FileService) GetLineCount(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	// Scan each line
	for scanner.Scan() {
		lineCount++
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	return lineCount, nil
}

func (fs FileService) GetWordCount(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a scanner that splits by words
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	return wordCount, nil
}