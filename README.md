# go2md: Go Documentation to Markdown Converter

go2md is a tool that reads Go source files and generates a corresponding markdown documentation based on GoDoc styled comments present in the code. It supports reading from a single file, from piped input, or recursively from a directory.

## Installation
Ensure you have Go installed on your system.

Navigate to the directory where go2md resides.

Run the following command to compile the program:

``go
go build
```

This will produce an executable named go2md in the current directory.

Usage
1. Recursively process a directory:

```bash
./go2md -r /path/to/directory -o /path/to/output.md
```

This will process all .go files found in the specified directory (and its sub-directories) and generate a markdown documentation in the output file.

2. Process a single file:

```bash
./go2md -i filename.go -o /path/to/output.md
```

This will process the specified .go file and generate the markdown documentation in the output file.

3. Process piped input:

```bash
cat filename.go | ./go2md -o /path/to/output.md
```

This allows you to pipe the content of a .go file into go2md and get the markdown documentation in the specified output file.

## Documentation Formatting Guide
In your Go code, to get the best results, follow this commenting style:

```go
// Title: Your Title Here
// Description: A brief description of the function or type.
// Function: The function signature or type definition.
// CalledWith: How the function can be called (for functions only).
// ExpectedOutput: What to expect as output (for functions only).
func YourFunction() {...}
```

**Title:** A brief title for the function or type.
**Description:** A more detailed description.
**Function:** The function signature or type definition.
**CalledWith:** (For functions) Example(s) of how the function can be called.
**ExpectedOutput:** (For functions) What the function returns or the expected output.

go2md will parse these comments and generate a structured markdown documentation.