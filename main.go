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
	stdOutWriter := bufio.NewWriter(os.Stdout)

	readFromAndWriteTo(stdInReader, stdOutWriter, "<<TODO>>")
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
