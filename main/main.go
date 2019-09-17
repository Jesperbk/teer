package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/jesperbk/teer/options"
)

func main() {
	options := options.OutputOptions{}

	arg.MustParse(&options)

	readFromAndWriteTo(os.Stdin, os.Stdout, options)
}

func readFromAndWriteTo(reader io.Reader, writer io.Writer, options options.OutputOptions) {
	inputReader := io.TeeReader(reader, writer)

	outputFile, err := openFile(options.OutputPath, options.DoTruncate)
	abortIfErr(err, "Error while creating/accessing file '%s': %v\n", options.OutputPath, err)

	sendInputToOutput(inputReader, outputFile)
	abortIfErr(err, "Error while writing to file '%s': %v\n", options.OutputPath, err)
}

func openFile(filePath string, doTruncate bool) (*os.File, error) {
	fileFlags := getFileOpenFlags(doTruncate)
	return os.OpenFile(filePath, fileFlags, 0644)
}

func sendInputToOutput(inputReader io.Reader, outputFile io.Writer) error {
	lineScanner := bufio.NewScanner(inputReader)

	for lineScanner.Scan() {
		_, err := fmt.Fprintln(outputFile, lineScanner.Text())
		if err != nil {
			return err
		}
	}

	return nil
}

func getFileOpenFlags(doTruncate bool) int {
	fileFlags := os.O_CREATE | os.O_WRONLY
	if doTruncate {
		fileFlags = fileFlags | os.O_TRUNC
	} else {
		fileFlags = fileFlags | os.O_APPEND
	}

	return fileFlags
}

func abortIfErr(err error, message string, args ...interface{}) {
	if err != nil {
		log.Fatalf(message, args...)
	}
}
