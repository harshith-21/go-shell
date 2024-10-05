package main

import (
    "bufio"
    "fmt"
    "os"
	"time"
	"strings"
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
    validCommands := map[string]func([]string, *bufio.Writer){
		"exit": exitCommand,      
    }

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

		// Split input into command and arguments
		parts := strings.Fields(input) // This will split the input by spaces
		if len(parts) == 0 {
			continue // Ignore empty input
		}

		command := parts[0]
		args := parts[1:]

		// Check if command is valid and execute it
		if cmdFunc, exists := validCommands[command]; exists {
			cmdFunc(args, writer) // Call the command function
		} else {
			fmt.Fprintln(writer, command, "Command not defined.")
		}

    }
}

func exitCommand(args []string, writer *bufio.Writer) {
    fmt.Fprintln(writer, "Exiting the Go Shell. Goodbye!")
    os.Exit(0) // Use os.Exit to terminate the program
}