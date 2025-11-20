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
	data, err := os.ReadFile(suite.tempFile.Name())
	suite.Require().NoError(err)
	size := suite.fs.GetFileSize(data)
	suite.Equal(len(suite.fileContent), size, "File size should match the content length")
}

// TestGetLineCountSingleLine tests getting the line count of a file with a single line.
func (suite *FileServiceTestSuite) TestGetLineCountSingleLine() {
	data, err := os.ReadFile(suite.tempFile.Name())
	suite.Require().NoError(err)
	count := suite.fs.GetLineCount(data)
	suite.Equal(0, count, "Line count should be 0 for single line file without newline")
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

	data, err := os.ReadFile(tmpfile.Name())
	suite.Require().NoError(err)
	count := suite.fs.GetLineCount(data)
	suite.Equal(2, count, "Line count should be 2 (newlines)")
}

// TestGetLineCountEmptyFile tests getting the line count of an empty file.
func (suite *FileServiceTestSuite) TestGetLineCountEmptyFile() {
	tmpfile, err := os.CreateTemp("", "empty")
	suite.Require().NoError(err)
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	data, err := os.ReadFile(tmpfile.Name())
	suite.Require().NoError(err)
	count := suite.fs.GetLineCount(data)
	suite.Equal(0, count, "Line count should be 0 for empty file")
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

	data, err := os.ReadFile(tmpfile.Name())
	suite.Require().NoError(err)
	count := suite.fs.GetWordCount(data)
	suite.Equal(1, count, "Word count should be 1")
}

// TestGetWordCountMultipleWords tests getting the word count of a file with multiple words.
func (suite *FileServiceTestSuite) TestGetWordCountMultipleWords() {
	// "hello world" in suite.tempFile
	data, err := os.ReadFile(suite.tempFile.Name())
	suite.Require().NoError(err)
	count := suite.fs.GetWordCount(data)
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

	data, err := os.ReadFile(tmpfile.Name())
	suite.Require().NoError(err)
	count := suite.fs.GetWordCount(data)
	suite.Equal(3, count, "Word count should be 3")
}

// TestGetWordCountEmptyFile tests getting the word count of an empty file.
func (suite *FileServiceTestSuite) TestGetWordCountEmptyFile() {
	tmpfile, err := os.CreateTemp("", "empty_words")
	suite.Require().NoError(err)
	defer os.Remove(tmpfile.Name())
	suite.Require().NoError(tmpfile.Close())

	data, err := os.ReadFile(tmpfile.Name())
	suite.Require().NoError(err)
	count := suite.fs.GetWordCount(data)
	suite.Equal(0, count, "Word count should be 0 for empty file")
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

	data, err := os.ReadFile(tmpfile.Name())
	suite.Require().NoError(err)
	count := suite.fs.GetCharCount(data)
	suite.Equal(1, count, "Char count should be 1")
}

// TestGetCharCountMultipleChars tests getting the char count of a file with multiple characters.
func (suite *FileServiceTestSuite) TestGetCharCountMultipleChars() {
	// "hello world" is 11 characters
	data, err := os.ReadFile(suite.tempFile.Name())
	suite.Require().NoError(err)
	count := suite.fs.GetCharCount(data)
	suite.Equal(11, count, "Char count should be 11 for 'hello world'")
}

// TestGetCharCountMultibyteChars tests getting the char count of a file with multibyte characters.
func (suite *FileServiceTestSuite) TestGetCharCountMultibyteChars() {
	content := "Hello 🌍"
	
	tmpfile, err := os.CreateTemp("", "multibyte")
	suite.Require().NoError(err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.WriteString(content)
	suite.Require().NoError(err)
	suite.Require().NoError(tmpfile.Close())

	data, err := os.ReadFile(tmpfile.Name())
	suite.Require().NoError(err)
	count := suite.fs.GetCharCount(data)
	suite.Equal(7, count, "Char count should be 7 for 'Hello 🌍'")
}

// TestGetCharCountEmptyFile tests getting the char count of an empty file.
func (suite *FileServiceTestSuite) TestGetCharCountEmptyFile() {
	tmpfile, err := os.CreateTemp("", "empty_chars")
	suite.Require().NoError(err)
	defer os.Remove(tmpfile.Name())
	suite.Require().NoError(tmpfile.Close())

	data, err := os.ReadFile(tmpfile.Name())
	suite.Require().NoError(err)
	count := suite.fs.GetCharCount(data)
	suite.Equal(0, count, "Char count should be 0 for empty file")
}

// TestFileServiceTestSuite runs the FileService test suite.
func TestFileServiceTestSuite(t *testing.T) {
	suite.Run(t, new(FileServiceTestSuite))
}
