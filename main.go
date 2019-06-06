package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	stdInReader := bufio.NewReader(os.Stdin)
	stdOutWriter := bufio.NewWriter(os.Stdout)

	readFromAndWriteTo(stdInReader, stdOutWriter)
}

func readFromAndWriteTo(reader io.Reader, writer io.Writer) {
	inputReader := io.TeeReader(reader, writer)
	lineScanner := bufio.NewScanner(inputReader)

	for lineScanner.Scan() {
	}
}
