package main

import (
    "bufio"
    "fmt"
    "os"
	"time"
	"strings"
	"sort"
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
		"ls": lsCommand,
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


func lsCommand(args []string, writer *bufio.Writer) {
    // Get the current working directory
    pwd, err := os.Getwd()
    if err != nil {
        fmt.Fprintln(writer, "Error getting current directory:", err)
        return
    }

    // Read the directory contents
    files, err := os.ReadDir(pwd)
    if err != nil {
        fmt.Fprintln(writer, "Error reading directory:", err)
        return
    }

    // Determine if hidden files should be included
    showHidden := false
    if len(args) > 0 && args[0] == "-a" {
        showHidden = true
    }

    // Collect file info and filter based on hidden file logic
    var fileInfos []os.DirEntry
    for _, file := range files {
        // Check if the file is hidden (starts with ".")
        if !showHidden && strings.HasPrefix(file.Name(), ".") {
            continue // Skip hidden files unless -a is provided
        }
        fileInfos = append(fileInfos, file)
    }

    // Sort files by modification time in ascending order
    sort.Slice(fileInfos, func(i, j int) bool {
        infoI, errI := fileInfos[i].Info()
        infoJ, errJ := fileInfos[j].Info()

        if errI != nil || errJ != nil {
            return false
        }

        return infoI.ModTime().Before(infoJ.ModTime())
    })

    // Loop through the sorted files and print their names with the last modified time and size
    for _, file := range fileInfos {
        // Get file info for each file
        info, err := file.Info()
        if err != nil {
            fmt.Fprintln(writer, "Error getting file info:", err)
            continue
        }

        // Get the last modified time
        modTime := info.ModTime().Format("2006-01-02 15:04:05")

        // Get the size of the file or directory
        size := formatSize(info.Size())

        // Print the last modified time, file size, and the file name
        fmt.Fprintf(writer, "%s  %10s  %s\n", modTime, size, file.Name())
    }

    writer.Flush() // Flush the writer to ensure output is shown
}


func formatSize(size int64) string {
    if size < 1024 {
        return fmt.Sprintf("%d B", size)
    } else if size < 1024*1024 {
        return fmt.Sprintf("%.2f KB", float64(size)/1024)
    } else if size < 1024*1024*1024 {
        return fmt.Sprintf("%.2f MB", float64(size)/(1024*1024))
    } else {
        return fmt.Sprintf("%.2f GB", float64(size)/(1024*1024*1024))
    }
}