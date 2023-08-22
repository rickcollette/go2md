// text2md: A tool to convert GoDoc style documentation to Markdown format.
//
// This utility can take input from a piped command, an input file or directly from the terminal.
// The output can be directed to a file or printed on the terminal.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// main is the entry point of the application.
func main() {
	// Define and parse command line flags
	input := flag.String("i", "", "Input file name. Read from the file instead of stdin.")
	output := flag.String("o", "", "Output file name. Write to the file instead of stdout.")

	flag.Parse()

	var reader io.Reader

	// Check if data is being piped in
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Reading from stdin
		reader = os.Stdin
	} else if *input != "" {
		// Reading from the specified input file
		file, err := os.Open(*input)
		if err != nil {
			fmt.Println("Error opening input file:", err)
			return
		}
		defer file.Close()
		reader = file
	} else {
		// No input source provided
		fmt.Println("No input provided!")
		return
	}

	var mdBuilder strings.Builder
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "    ") {
			// Convert indented lines to code blocks in Markdown
			mdBuilder.WriteString("```go\n")
			mdBuilder.WriteString(strings.TrimPrefix(line, "    ") + "\n")
			mdBuilder.WriteString("```\n")
		} else if strings.HasPrefix(line, "TYPE") || strings.HasPrefix(line, "FUNC") {
			// Convert TYPE or FUNC prefixes to headers in Markdown
			mdBuilder.WriteString("## " + line + "\n")
		} else {
			// Regular lines are preserved as they are
			mdBuilder.WriteString(line + "\n")
		}
	}

	// Determine output destination and write the generated markdown content
	if *output != "" {
		// Write to the specified output file
		err := os.WriteFile(*output, []byte(mdBuilder.String()), 0644)
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
		fmt.Printf("Documentation saved to %s!\n", *output)
	} else {
		// Print to stdout
		fmt.Print(mdBuilder.String())
	}
}
