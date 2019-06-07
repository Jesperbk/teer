package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

var input string

func TestWriteToStdOut(t *testing.T) {
	input = "Some string\nSome other string"
	inputReader := readerOf(input)
	var outputBuffer bytes.Buffer

	testFilePath, testDirPath := createTestFile(t)
	defer os.RemoveAll(testDirPath)

	readFromAndWriteTo(inputReader, &outputBuffer, testFilePath)

	validateOutputFromReader(t, &outputBuffer)
}

func TestWriteToFile(t *testing.T) {
	input = "Some string\nSome other string"
	inputReader := readerOf(input)

	testFilePath, testDirPath := createTestFile(t)
	defer os.RemoveAll(testDirPath)

	readFromAndWriteTo(inputReader, ioutil.Discard, testFilePath)

	validateOutputFromFilePath(t, testFilePath)
}

func readerOf(str string) io.Reader {
	return strings.NewReader(str)
}

func createTestFile(t *testing.T) (string, string) {
	testDirPath, err := ioutil.TempDir("", "teer_test")
	if err != nil {
		t.Fatal(err)
	}
	testFilePath := path.Join(testDirPath, "test_file.log")

	return testFilePath, testDirPath
}

func validateOutputFromReader(t *testing.T, reader io.Reader) {
	outputBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatalf("Error while reading output: %v", err)
	}
	output := string(outputBytes)

	validateOutput(t, output)
}

func validateOutputFromFilePath(t *testing.T, path string) {
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Error while opening test file '%s': %v", path, err)
	}

	validateOutputFromReader(t, file)
}

func validateOutput(t *testing.T, output string) {
	if stripTrailingNewLine(output) != input {
		t.Fatalf("Unexpected output: '%s'\n", output)
	}
}

func stripTrailingNewLine(str string) string {
	return strings.TrimSuffix(str, "\n")
}
