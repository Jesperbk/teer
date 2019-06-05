package main

import (
	"bufio"
	"os"
)

func main() {
	stdInReader := bufio.NewReader(os.Stdin)
	stdOutWriter := bufio.NewWriter(os.Stdout)

	readFromAndWriteTo(stdInReader, stdOutWriter)
}

func readFromAndWriteTo(reader *bufio.Reader, writer *bufio.Writer) {

}
