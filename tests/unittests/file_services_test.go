package tests

import (
	"strings"
	"testing"

	"github.com/abeni-al7/aben-wc/services"
	"github.com/stretchr/testify/suite"
)

// FileServiceTestSuite is a test suite for the FileService.
type FileServiceTestSuite struct {
	suite.Suite
	fs services.FileService
}

// SetupTest creates a temporary file and directory for testing.
func (suite *FileServiceTestSuite) SetupTest() {
	suite.fs = services.FileService{}
}

func (suite *FileServiceTestSuite) TestCalculateCounts_EmptyFile() {
	input := ""
	expected := services.Counts{
		Bytes: 0,
		Lines: 0,
		Words: 0,
		Chars: 0,
	}
	r := strings.NewReader(input)
	counts, err := suite.fs.CalculateCounts(r)
	suite.Require().NoError(err)
	suite.Equal(expected, counts)
}

func (suite *FileServiceTestSuite) TestCalculateCounts_SingleWord() {
	input := "hello"
	expected := services.Counts{
		Bytes: 5,
		Lines: 0,
		Words: 1,
		Chars: 5,
	}
	r := strings.NewReader(input)
	counts, err := suite.fs.CalculateCounts(r)
	suite.Require().NoError(err)
	suite.Equal(expected, counts)
}

func (suite *FileServiceTestSuite) TestCalculateCounts_MultipleWords() {
	input := "hello world"
	expected := services.Counts{
		Bytes: 11,
		Lines: 0,
		Words: 2,
		Chars: 11,
	}
	r := strings.NewReader(input)
	counts, err := suite.fs.CalculateCounts(r)
	suite.Require().NoError(err)
	suite.Equal(expected, counts)
}

func (suite *FileServiceTestSuite) TestCalculateCounts_MultipleLines() {
	input := "line1\nline2\nline3"
	expected := services.Counts{
		Bytes: 17,
		Lines: 2,
		Words: 3,
		Chars: 17,
	}
	r := strings.NewReader(input)
	counts, err := suite.fs.CalculateCounts(r)
	suite.Require().NoError(err)
	suite.Equal(expected, counts)
}

func (suite *FileServiceTestSuite) TestCalculateCounts_MultibyteCharacters() {
	input := "Hello 🌍"
	expected := services.Counts{
		Bytes: 10, // Hello (5) + space (1) + 🌍 (4)
		Lines: 0,
		Words: 2,
		Chars: 7, // Hello (5) + space (1) + 🌍 (1)
	}
	r := strings.NewReader(input)
	counts, err := suite.fs.CalculateCounts(r)
	suite.Require().NoError(err)
	suite.Equal(expected, counts)
}

func (suite *FileServiceTestSuite) TestCalculateCounts_TrailingNewline() {
	input := "hello\n"
	expected := services.Counts{
		Bytes: 6,
		Lines: 1,
		Words: 1,
		Chars: 6,
	}
	r := strings.NewReader(input)
	counts, err := suite.fs.CalculateCounts(r)
	suite.Require().NoError(err)
	suite.Equal(expected, counts)
}

func (suite *FileServiceTestSuite) TestCalculateCounts_MultipleSpaces() {
	input := "hello   world"
	expected := services.Counts{
		Bytes: 13,
		Lines: 0,
		Words: 2,
		Chars: 13,
	}
	r := strings.NewReader(input)
	counts, err := suite.fs.CalculateCounts(r)
	suite.Require().NoError(err)
	suite.Equal(expected, counts)
}

// TestFileServiceTestSuite runs the FileService test suite.
func TestFileServiceTestSuite(t *testing.T) {
	suite.Run(t, new(FileServiceTestSuite))
}
