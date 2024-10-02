package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {

	// Create a buffered reader to read from standard input
	reader := bufio.NewReader(os.Stdin)

	// Create a buffered writer to write to standard output
	writer := bufio.NewWriter(os.Stdout)

	// Remember to flush the writer after writing
	defer writer.Flush()
	
	// Start the REPL loop (we'll implement this next)
	repl(reader, writer)

}

func repl(reader *bufio.Reader, writer *bufio.Writer) {
    // Placeholder for the REPL functionality
    fmt.Fprintln(writer, "Welcome to the Go Shell!")
}