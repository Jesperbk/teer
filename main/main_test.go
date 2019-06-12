package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/jesperbk/teer/options"
)

func TestWriteToStdOut(t *testing.T) {
	input := "Some string\nSome other string"
	inputReader := readerOf(input)
	var outputBuffer bytes.Buffer

	testFilePath, testDirPath := createTestFilePath(t)
	defer os.RemoveAll(testDirPath)

	options := options.OutputOptions{
		OutputPath:   testFilePath,
		ExistsAction: options.ExistsActionTruncate,
	}
	readFromAndWriteTo(inputReader, &outputBuffer, options)

	expected := input
	validateOutputFromReader(t, expected, &outputBuffer)
}

func TestWriteToFile(t *testing.T) {
	input := "Some string\nSome other string"
	inputReader := readerOf(input)

	testFilePath, testDirPath := createTestFilePath(t)
	defer os.RemoveAll(testDirPath)

	options := options.OutputOptions{
		OutputPath:   testFilePath,
		ExistsAction: options.ExistsActionTruncate,
	}
	readFromAndWriteTo(inputReader, ioutil.Discard, options)

	expected := input
	validateOutputFromFilePath(t, expected, testFilePath)
}

func TestCreateWhenFileNotExisting(t *testing.T) {
	input := "Some string\nSome other string"
	inputReader := readerOf(input)

	testFilePath, testDirPath := createTestFilePath(t)
	defer os.RemoveAll(testDirPath)

	options := options.OutputOptions{
		OutputPath:   testFilePath,
		ExistsAction: options.ExistsActionAppend,
	}
	readFromAndWriteTo(inputReader, ioutil.Discard, options)

	expected := input
	validateOutputFromFilePath(t, expected, testFilePath)
}

func TestOverwriteExistingFile(t *testing.T) {
	input := "Some string\nSome other string"
	inputReader := readerOf(input)

	testFilePath, testDirPath := createTestFileWithContent(t, "Old content")
	defer os.RemoveAll(testDirPath)

	options := options.OutputOptions{
		OutputPath:   testFilePath,
		ExistsAction: options.ExistsActionTruncate,
	}
	readFromAndWriteTo(inputReader, ioutil.Discard, options)

	expected := input
	validateOutputFromFilePath(t, expected, testFilePath)
}

func TestAppendToExistingFile(t *testing.T) {
	input := "Some string\nSome other string"
	inputReader := readerOf(input)

	testFilePath, testDirPath := createTestFileWithContent(t, "Old content")
	defer os.RemoveAll(testDirPath)

	options := options.OutputOptions{
		OutputPath:   testFilePath,
		ExistsAction: options.ExistsActionAppend,
	}
	readFromAndWriteTo(inputReader, ioutil.Discard, options)

	expected := "Old content\n" + input
	validateOutputFromFilePath(t, expected, testFilePath)
}

func readerOf(str string) io.Reader {
	return strings.NewReader(str)
}

func createTestFilePath(t *testing.T) (string, string) {
	testDirPath, err := ioutil.TempDir("", "teer_test")
	if err != nil {
		t.Fatal(err)
	}
	testFilePath := path.Join(testDirPath, "test_file.log")

	return testFilePath, testDirPath
}

func createTestFileWithContent(t *testing.T, content string) (string, string) {
	testFilePath, testDirPath := createTestFilePath(t)

	testFile, err := os.Create(testFilePath)
	defer testFile.Close()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Fprintln(testFile, content)

	return testFilePath, testDirPath
}

func validateOutputFromReader(t *testing.T, expected string, reader io.Reader) {
	outputBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatalf("Error while reading output: %v", err)
	}
	output := string(outputBytes)

	validateOutput(t, expected, output)
}

func validateOutputFromFilePath(t *testing.T, expected string, path string) {
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Error while opening test file '%s': %v", path, err)
	}

	validateOutputFromReader(t, expected, file)
}

func validateOutput(t *testing.T, expected string, output string) {
	if stripTrailingNewLine(output) != expected {
		t.Fatalf("Unexpected output: '%s'\n", output)
	}
}

func stripTrailingNewLine(str string) string {
	return strings.TrimSuffix(str, "\n")
}
