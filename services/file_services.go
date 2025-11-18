package services

import (
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
