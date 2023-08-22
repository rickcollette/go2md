# go2md
## Go Source to Markdown Documentation Creator

go2md is a tool that reads Go source files and generates a corresponding markdown documentation based on go2md styled comments present in the code. It supports reading from a single file, from piped input, or recursively from a directory.

## Installation
Ensure you have Go installed on your system.

Navigate to the directory where go2md resides.

Run the following command to compile the program:

```go
go build -o go2md main.go
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

For the top of the file:

```go
// Package: main
// Description: Description of the package and its purpose.
// Git Repository: URL of the git repository.
// License: License type of the code.
```

For functions and types:

```go
// Title: Your Title Here
// Description: A brief description of the function or type.
// Function: The function signature or type definition.
// CalledWith: How the function can be called (for functions only).
// ExpectedOutput: What to expect as output (for functions only).
// Example: An example of how to use the function or type.
func YourFunction() {...}
```

* **Title:** A brief title for the function or type.
* **Description:** A more detailed description.
* **Function:** The function signature or type definition.
* **CalledWith:** (For functions) Example(s) of how the function can be called.
* **ExpectedOutput:** (For functions) What the function returns or the expected output.
* **Example:** (For functions and types) An example showcasing the usage.


go2md will parse these comments and generate a structured markdown documentation.

-------------------

## Generated with go2md:

# Package:  main

**Description:**  converts go code into documentation

**Git Repository:**  https://github.com/rickcollette/go2md

**License:**  MIT License

## Function: main

**Title:**  main

**Description:**  Entry point of the program. Converts Go source documentation to Markdown format based on provided flags.

**Function:**  main

**Called With:**  N/A

**Example:**  N/A (Since it's the main function)

**Expected Output:**  Writes markdown content to a file or stdout.

## Function: processFile

**Title:**  processFile

**Description:**  Process a Go file and generate markdown documentation.

**Function:**  processFile

**Called With:**  filename string, mdBuilder *strings.Builder

**Example:**  processFile("example.go", &mdBuilder)

**Expected Output:**  Markdown documentation for the Go file.

## Function: processGenDecl

**Title:**  processGenDecl

**Description:**  Process general declarations (like variables and types) and generates markdown documentation.

**Function:**  processGenDecl

**Called With:**  decl *ast.GenDecl, mdBuilder *strings.Builder

**Example:**  processGenDecl(someGenDecl, &mdBuilder)

**Expected Output:**  Markdown documentation for general declarations.

## Function: processFuncDecl

**Title:**  processFuncDecl

**Description:**  Process function declarations and generates markdown documentation based on specific comment patterns.

**Function:**  processFuncDecl

**Called With:**  decl *ast.FuncDecl, mdBuilder *strings.Builder

**Example:**  processFuncDecl(someFuncDecl, &mdBuilder)

**Expected Output:**  Markdown documentation for function declarations.

