package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run . <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// *os.File implements io.Reader, and os.Stdout implements io.Writer —
	// so io.Copy can stream the file straight to the terminal for us,
	// without a manual read-loop into a byte buffer.
	if _, err := io.Copy(os.Stdout, file); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
}
