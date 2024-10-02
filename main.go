package main

import (
    "bufio"
    "fmt"
    "os"
	"time"
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

	// Placeholder for valid commands
    validCommands := []string{}

	for {
        // Print the prompt (pwd will be added later)
		currentTime := time.Now().Format("15:04:05")
        fmt.Fprintf(writer, "%s > ", currentTime)
        writer.Flush()

        // Read user input
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(writer, "Error reading input:", err)
            continue
        }

        // Remove newline character from input
        input = input[:len(input)-1]

        // Check for exit command
        if input == "exit" {
            fmt.Fprintln(writer, "Exiting the Go Shell. Goodbye!")
            break
        }

        // Validate command
        isValid := false
        for _, command := range validCommands {
            if input == command {
                isValid = true
                break
            }
        }

        // Respond based on command validity
        if isValid {
            fmt.Fprintf(writer, "Executing command: %s\n", input)
        } else {
            fmt.Fprintln(writer, string(input) + " command not defined.")
        }
    }
}