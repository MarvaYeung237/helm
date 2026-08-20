package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

// ReadValues reads from an io.Reader and handles the EOF case correctly
// to ensure no data is lost if the input ends without a newline.
func ReadValues(r io.Reader) ([]byte, error) {
	var buf bytes.Buffer
	reader := bufio.NewReader(r)
	for {
		line, err := reader.ReadBytes('\n')
		buf.Write(line)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func main() {
	data, err := ReadValues(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading values: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Read %d bytes\n", len(data))
}