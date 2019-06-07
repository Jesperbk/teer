package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	stdInReader := bufio.NewReader(os.Stdin)

	readFromAndWriteTo(stdInReader, os.Stdout, os.Args[1])
}

func readFromAndWriteTo(reader io.Reader, writer io.Writer, outputFilePath string) {
	inputReader := io.TeeReader(reader, writer)
	lineScanner := bufio.NewScanner(inputReader)

	outputFile, err := os.Create(outputFilePath)
	abortIfErr(err, "Error while creating/accessing file '%s': %v\n", outputFilePath, err)

	for lineScanner.Scan() {
		_, err := fmt.Fprintln(outputFile, lineScanner.Text())
		abortIfErr(err, "Error while writing to file '%s': %v\n", outputFilePath, err)
	}
}

func abortIfErr(err error, message string, args ...interface{}) {
	if err != nil {
		log.Fatalf(message, args...)
	}
}
