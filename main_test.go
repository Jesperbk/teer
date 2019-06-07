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
)

func TestWriteToStdOut(t *testing.T) {
	input := "Some string\nSome other string"
	inputReader := strings.NewReader(input)
	var outputBuffer bytes.Buffer

	testDir, err := ioutil.TempDir("", "teer_test")
	testFilePath := path.Join(testDir, "test_file.log")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(testDir)

	readFromAndWriteTo(inputReader, &outputBuffer, testFilePath)

	output, err := getOutput(&outputBuffer)
	if err != nil {
		t.Fatal(err)
	}
	if input != output {
		t.Fatalf("Unexpected output: '%s'\n", output)
	}
}

func TestWriteToFile(t *testing.T) {
	input := "Some string\nSome other string"
	inputReader := strings.NewReader(input)

	testDir, err := ioutil.TempDir("", "teer_test")
	testFilePath := path.Join(testDir, "test_file.log")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(testDir)

	readFromAndWriteTo(inputReader, ioutil.Discard, testFilePath)

	outputBytes, err := ioutil.ReadFile(testFilePath)
	output := string(outputBytes)
	if input != output {
		t.Fatalf("Unexpected output: '%s'\n", output)
	}
}

func getOutput(reader io.Reader) (string, error) {
	outputBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("Error while reading output: %v", err)
	}
	output := string(outputBytes)
	return output, nil
}
