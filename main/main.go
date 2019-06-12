package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jesperbk/teer/options"
)

func main() {
	options := options.OutputOptions{
		OutputPath:   os.Args[1],
		ExistsAction: options.ExistsActionTruncate,
	}

	readFromAndWriteTo(os.Stdin, os.Stdout, options)
}

func readFromAndWriteTo(reader io.Reader, writer io.Writer, options options.OutputOptions) {
	inputReader := io.TeeReader(reader, writer)

	outputFile, err := openFile(options.OutputPath, options.ExistsAction)
	abortIfErr(err, "Error while creating/accessing file '%s': %v\n", options.OutputPath, err)

	sendInputToOutput(inputReader, outputFile)
	abortIfErr(err, "Error while writing to file '%s': %v\n", options.OutputPath, err)
}

func openFile(filePath string, existsAction options.ExistsAction) (*os.File, error) {
	fileFlags := getFileOpenFlags(existsAction)
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

func getFileOpenFlags(existsAction options.ExistsAction) int {
	fileFlags := os.O_CREATE | os.O_WRONLY
	switch existsAction {
	case options.ExistsActionTruncate:
		fileFlags = fileFlags | os.O_TRUNC
	case options.ExistsActionAppend:
		fileFlags = fileFlags | os.O_APPEND
	default:
		log.Fatalf("Internal error! Attempting to resolve invalid WriteMode: %v\n", existsAction)
	}

	return fileFlags
}

func abortIfErr(err error, message string, args ...interface{}) {
	if err != nil {
		log.Fatalf(message, args...)
	}
}
