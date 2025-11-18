package tests

import (
	"os"
	"testing"

	"github.com/abeni-al7/aben-wc/services"
	"github.com/stretchr/testify/suite"
)

// FileServiceTestSuite is a test suite for the FileService.
type FileServiceTestSuite struct {
	suite.Suite
	fs          services.FileService
	tempFile    *os.File
	tempDir     string
	fileContent string
}

// SetupTest creates a temporary file and directory for testing.
func (suite *FileServiceTestSuite) SetupTest() {
	suite.fs = services.FileService{}
	suite.fileContent = "hello world"

	// Create a temporary file with some content
	tmpfile, err := os.CreateTemp("", "example")
	suite.Require().NoError(err, "Failed to create temp file")
	_, err = tmpfile.Write([]byte(suite.fileContent))
	suite.Require().NoError(err, "Failed to write to temp file")
	err = tmpfile.Close()
	suite.Require().NoError(err, "Failed to close temp file")
	suite.tempFile = tmpfile

	// Create a temporary directory
	tmpdir, err := os.MkdirTemp("", "exampledir")
	suite.Require().NoError(err, "Failed to create temp dir")
	suite.tempDir = tmpdir
}

// TearDownTest removes the temporary file and directory.
func (suite *FileServiceTestSuite) TearDownTest() {
	err := os.Remove(suite.tempFile.Name())
	suite.NoError(err, "Failed to remove temp file")
	err = os.RemoveAll(suite.tempDir)
	suite.NoError(err, "Failed to remove temp dir")
}

// TestGetFileSizeExistingFile tests getting the size of an existing file.
func (suite *FileServiceTestSuite) TestGetFileSizeExistingFile() {
	size, err := suite.fs.GetFileSize(suite.tempFile.Name())
	suite.NoError(err, "Should not return an error for an existing file")
	suite.Equal(int64(len(suite.fileContent)), size, "File size should match the content length")
}

// TestGetFileSizeNonExistentFile tests getting the size of a non-existent file.
func (suite *FileServiceTestSuite) TestGetFileSizeNonExistentFile() {
	_, err := suite.fs.GetFileSize("non-existent-file.txt")
	suite.Error(err, "Should return an error for a non-existent file")
}

// TestGetFileSizeDirectory tests getting the size of a directory.
func (suite *FileServiceTestSuite) TestGetFileSizeDirectory() {
	_, err := suite.fs.GetFileSize(suite.tempDir)
	suite.Error(err, "Should return an error for a directory")
}

// TestFileServiceTestSuite runs the FileService test suite.
func TestFileServiceTestSuite(t *testing.T) {
	suite.Run(t, new(FileServiceTestSuite))
}
