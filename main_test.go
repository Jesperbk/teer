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

func TestWriteToStdOut(t *testing.T) {
	input := "Some string\nSome other string"
	inputReader := readerOf(input)
	var outputBuffer bytes.Buffer

	testFilePath, testDirPath := createTestFile(t)
	defer os.RemoveAll(testDirPath)

	readFromAndWriteTo(inputReader, &outputBuffer, testFilePath)

	output := readAllFromReader(t, &outputBuffer)
	if !matches(output, input) {
		t.Fatalf("Unexpected output: '%s'\n", output)
	}
}

func TestWriteToFile(t *testing.T) {
	input := "Some string\nSome other string"
	inputReader := readerOf(input)

	testFilePath, testDirPath := createTestFile(t)
	defer os.RemoveAll(testDirPath)

	readFromAndWriteTo(inputReader, ioutil.Discard, testFilePath)

	output := readAllFromFilePath(t, testFilePath)
	if !matches(output, input) {
		t.Fatalf("Unexpected output: '%s'\n", output)
	}
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

func readAllFromReader(t *testing.T, reader io.Reader) string {
	outputBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatalf("Error while reading output: %v", err)
	}
	output := string(outputBytes)
	return output
}

func readAllFromFilePath(t *testing.T, path string) string {
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Error while opening test file '%s': %v", path, err)
	}
	return readAllFromReader(t, file)
}

func matches(output string, input string) bool {
	return input == stripTrailingNewLine(output)
}

func stripTrailingNewLine(str string) string {
	return strings.TrimSuffix(str, "\n")
}
