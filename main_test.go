package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestWriteToStdOut(t *testing.T) {
	input := "Some string\nSome other string"
	inputReader := strings.NewReader(input)
	var outputBuffer bytes.Buffer

	readFromAndWriteTo(inputReader, &outputBuffer)

	output, err := getOutput(&outputBuffer)
	if err != nil {
		t.Fatal(err)
	}
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
