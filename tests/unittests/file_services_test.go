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

// TestGetLineCountSingleLine tests getting the line count of a file with a single line.
func (suite *FileServiceTestSuite) TestGetLineCountSingleLine() {
	count, err := suite.fs.GetLineCount(suite.tempFile.Name())
	suite.NoError(err, "Should not return an error for an existing file")
	suite.Equal(1, count, "Line count should be 1 for single line file")
}

// TestGetLineCountMultipleLines tests getting the line count of a file with multiple lines.
func (suite *FileServiceTestSuite) TestGetLineCountMultipleLines() {
	content := "line1\nline2\nline3"
	tmpfile, err := os.CreateTemp("", "multiline")
	suite.Require().NoError(err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.WriteString(content)
	suite.Require().NoError(err)
	suite.Require().NoError(tmpfile.Close())

	count, err := suite.fs.GetLineCount(tmpfile.Name())
	suite.NoError(err)
	suite.Equal(3, count, "Line count should be 3")
}

// TestGetLineCountEmptyFile tests getting the line count of an empty file.
func (suite *FileServiceTestSuite) TestGetLineCountEmptyFile() {
	tmpfile, err := os.CreateTemp("", "empty")
	suite.Require().NoError(err)
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	count, err := suite.fs.GetLineCount(tmpfile.Name())
	suite.NoError(err)
	suite.Equal(0, count, "Line count should be 0 for empty file")
}

// TestGetLineCountNonExistentFile tests getting the line count of a non-existent file.
func (suite *FileServiceTestSuite) TestGetLineCountNonExistentFile() {
	_, err := suite.fs.GetLineCount("non-existent-file.txt")
	suite.Error(err, "Should return an error for a non-existent file")
}

// TestGetLineCountDirectory tests getting the line count of a directory.
func (suite *FileServiceTestSuite) TestGetLineCountDirectory() {
	_, err := suite.fs.GetLineCount(suite.tempDir)
	suite.Error(err, "Should return an error for a directory")
}

// TestGetWordCountSingleWord tests getting the word count of a file with a single word.
func (suite *FileServiceTestSuite) TestGetWordCountSingleWord() {
	content := "hello"
	tmpfile, err := os.CreateTemp("", "singleword")
	suite.Require().NoError(err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.WriteString(content)
	suite.Require().NoError(err)
	suite.Require().NoError(tmpfile.Close())

	count, err := suite.fs.GetWordCount(tmpfile.Name())
	suite.NoError(err)
	suite.Equal(1, count, "Word count should be 1")
}

// TestGetWordCountMultipleWords tests getting the word count of a file with multiple words.
func (suite *FileServiceTestSuite) TestGetWordCountMultipleWords() {
	// "hello world" in suite.tempFile
	count, err := suite.fs.GetWordCount(suite.tempFile.Name())
	suite.NoError(err)
	suite.Equal(2, count, "Word count should be 2 for 'hello world'")
}

// TestGetWordCountMultipleLines tests getting the word count of a file with words on multiple lines.
func (suite *FileServiceTestSuite) TestGetWordCountMultipleLines() {
	content := "hello\nworld\nagain"
	tmpfile, err := os.CreateTemp("", "multiline_words")
	suite.Require().NoError(err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.WriteString(content)
	suite.Require().NoError(err)
	suite.Require().NoError(tmpfile.Close())

	count, err := suite.fs.GetWordCount(tmpfile.Name())
	suite.NoError(err)
	suite.Equal(3, count, "Word count should be 3")
}

// TestGetWordCountEmptyFile tests getting the word count of an empty file.
func (suite *FileServiceTestSuite) TestGetWordCountEmptyFile() {
	tmpfile, err := os.CreateTemp("", "empty_words")
	suite.Require().NoError(err)
	defer os.Remove(tmpfile.Name())
	suite.Require().NoError(tmpfile.Close())

	count, err := suite.fs.GetWordCount(tmpfile.Name())
	suite.NoError(err)
	suite.Equal(0, count, "Word count should be 0 for empty file")
}

// TestGetWordCountNonExistentFile tests getting the word count of a non-existent file.
func (suite *FileServiceTestSuite) TestGetWordCountNonExistentFile() {
	_, err := suite.fs.GetWordCount("non-existent-file.txt")
	suite.Error(err, "Should return an error for a non-existent file")
}

// TestGetWordCountDirectory tests getting the word count of a directory.
func (suite *FileServiceTestSuite) TestGetWordCountDirectory() {
	_, err := suite.fs.GetWordCount(suite.tempDir)
	suite.Error(err, "Should return an error for a directory")
}

// TestGetCharCountSingleChar tests getting the char count of a file with a single character.
func (suite *FileServiceTestSuite) TestGetCharCountSingleChar() {
	content := "a"
	tmpfile, err := os.CreateTemp("", "singlechar")
	suite.Require().NoError(err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.WriteString(content)
	suite.Require().NoError(err)
	suite.Require().NoError(tmpfile.Close())

	count, err := suite.fs.GetCharCount(tmpfile.Name())
	suite.NoError(err)
	suite.Equal(1, count, "Char count should be 1")
}

// TestGetCharCountMultipleChars tests getting the char count of a file with multiple characters.
func (suite *FileServiceTestSuite) TestGetCharCountMultipleChars() {
	// "hello world" is 11 characters
	count, err := suite.fs.GetCharCount(suite.tempFile.Name())
	suite.NoError(err)
	suite.Equal(11, count, "Char count should be 11 for 'hello world'")
}

// TestGetCharCountMultibyteChars tests getting the char count of a file with multibyte characters.
func (suite *FileServiceTestSuite) TestGetCharCountMultibyteChars() {
	content := "Hello 🌍" // 6 chars + space + emoji = 7 runes
	
	tmpfile, err := os.CreateTemp("", "multibyte")
	suite.Require().NoError(err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.WriteString(content)
	suite.Require().NoError(err)
	suite.Require().NoError(tmpfile.Close())

	count, err := suite.fs.GetCharCount(tmpfile.Name())
	suite.NoError(err)
	suite.Equal(7, count, "Char count should be 7 for 'Hello 🌍'")
}

// TestGetCharCountEmptyFile tests getting the char count of an empty file.
func (suite *FileServiceTestSuite) TestGetCharCountEmptyFile() {
	tmpfile, err := os.CreateTemp("", "empty_chars")
	suite.Require().NoError(err)
	defer os.Remove(tmpfile.Name())
	suite.Require().NoError(tmpfile.Close())

	count, err := suite.fs.GetCharCount(tmpfile.Name())
	suite.NoError(err)
	suite.Equal(0, count, "Char count should be 0 for empty file")
}

// TestGetCharCountNonExistentFile tests getting the char count of a non-existent file.
func (suite *FileServiceTestSuite) TestGetCharCountNonExistentFile() {
	_, err := suite.fs.GetCharCount("non-existent-file.txt")
	suite.Error(err, "Should return an error for a non-existent file")
}

// TestGetCharCountDirectory tests getting the char count of a directory.
func (suite *FileServiceTestSuite) TestGetCharCountDirectory() {
	_, err := suite.fs.GetCharCount(suite.tempDir)
	suite.Error(err, "Should return an error for a directory")
}

// TestFileServiceTestSuite runs the FileService test suite.
func TestFileServiceTestSuite(t *testing.T) {
	suite.Run(t, new(FileServiceTestSuite))
}
