// Package: main
// Description: converts go code into documentation
// Git Repository: https://github.com/rickcollette/go2md
// License: MIT License
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// Title: main
// Description: Entry point of the program. Converts Go source documentation to Markdown format based on provided flags.
// Function: main
// CalledWith: N/A
// ExpectedOutput: Writes markdown content to a file or stdout.
// Example: N/A (Since it's the main function)
func main() {
	dir := flag.String("r", "", "Directory to scan for .go files.")
	inputFile := flag.String("i", "", "Input .go file.")
	output := flag.String("o", "", "Output file name. Write to the file instead of stdout.")
	flag.Parse()

	var mdBuilder strings.Builder

	if *dir != "" {
		// Recursively process directory
		filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, ".go") {
				processFile(path, &mdBuilder)
			}
			return nil
		})
	} else if *inputFile != "" {
		// Process single file
		processFile(*inputFile, &mdBuilder)
	}

	// Determine output destination and write the generated markdown content
	if *output != "" {
		err := os.WriteFile(*output, []byte(mdBuilder.String()), 0644)
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
		fmt.Printf("Documentation saved to %s!\n", *output)
	} else {
		fmt.Print(mdBuilder.String())
	}
}

// Title: processFile
// Description: Process a Go file and generate markdown documentation.
// Function: processFile
// CalledWith: filename string, mdBuilder *strings.Builder
// ExpectedOutput: Markdown documentation for the Go file.
// Example: processFile("example.go", &mdBuilder)
func processFile(filename string, mdBuilder *strings.Builder) {
    fileSet := token.NewFileSet()
    content, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("Error reading the file:", err)
        return
    }

    formattedSource, err := format.Source(content)
    if err != nil {
        fmt.Println("Error formatting the source:", err)
        return
    }

    node, err := parser.ParseFile(fileSet, filename, formattedSource, parser.ParseComments)
    if err != nil {
        fmt.Println("Error parsing Go file:", err)
        return
    }

    // Process top-of-file package-level documentation
    if node.Doc != nil {
        var packageName, description, gitRepo, license string
        for _, comment := range node.Doc.List {
            if strings.HasPrefix(comment.Text, "// Package:") {
                packageName = strings.TrimPrefix(comment.Text, "// Package:")
            } else if strings.HasPrefix(comment.Text, "// Description:") {
                description = strings.TrimPrefix(comment.Text, "// Description:")
            } else if strings.HasPrefix(comment.Text, "// Git Repository:") {
                gitRepo = strings.TrimPrefix(comment.Text, "// Git Repository:")
            } else if strings.HasPrefix(comment.Text, "// License:") {
                license = strings.TrimPrefix(comment.Text, "// License:")
            } else {
                // For other comments that don't match the specific tags, append them as-is
                mdBuilder.WriteString(comment.Text + "\n")
            }
        }
        mdBuilder.WriteString(fmt.Sprintf("# Package: %s\n\n**Description:** %s\n\n**Git Repository:** %s\n\n**License:** %s\n\n",
                                          packageName, description, gitRepo, license))
    }

    for _, decl := range node.Decls {
        switch d := decl.(type) {
        case *ast.GenDecl:
            processGenDecl(d, mdBuilder)
        case *ast.FuncDecl:
            processFuncDecl(d, mdBuilder)
        }
    }
}


// Title: processGenDecl
// Description: Process general declarations (like variables and types) and generates markdown documentation.
// Function: processGenDecl
// CalledWith: decl *ast.GenDecl, mdBuilder *strings.Builder
// ExpectedOutput: Markdown documentation for general declarations.
// Example: processGenDecl(someGenDecl, &mdBuilder)
func processGenDecl(decl *ast.GenDecl, mdBuilder *strings.Builder) {
	for _, spec := range decl.Specs {
		switch s := spec.(type) {
		case *ast.ValueSpec:
			for _, name := range s.Names {
				mdBuilder.WriteString(fmt.Sprintf("## %s (%s)\n\n", name.Name, decl.Tok))
				if s.Doc != nil {
					mdBuilder.WriteString(s.Doc.Text() + "\n\n")
				}
			}
		case *ast.TypeSpec:
			mdBuilder.WriteString(fmt.Sprintf("## Type: %s\n\n", s.Name.Name))
			if s.Doc != nil {
				mdBuilder.WriteString(s.Doc.Text() + "\n\n")
			}
		}
	}
}

// Title: processFuncDecl
// Description: Process function declarations and generates markdown documentation based on specific comment patterns.
// Function: processFuncDecl
// CalledWith: decl *ast.FuncDecl, mdBuilder *strings.Builder
// ExpectedOutput: Markdown documentation for function declarations.
// Example: processFuncDecl(someFuncDecl, &mdBuilder)
func processFuncDecl(decl *ast.FuncDecl, mdBuilder *strings.Builder) {
	mdBuilder.WriteString(fmt.Sprintf("## Function: %s\n\n", decl.Name.Name))
	if decl.Doc != nil {
		var title, description, function, calledWith, example, expectedOutput string
		for _, comment := range decl.Doc.List {
			if strings.HasPrefix(comment.Text, "// Title:") {
				title = strings.TrimPrefix(comment.Text, "// Title:")
			} else if strings.HasPrefix(comment.Text, "// Description:") {
				description = strings.TrimPrefix(comment.Text, "// Description:")
			} else if strings.HasPrefix(comment.Text, "// Function:") {
				function = strings.TrimPrefix(comment.Text, "// Function:")
			} else if strings.HasPrefix(comment.Text, "// CalledWith:") {
				calledWith = strings.TrimPrefix(comment.Text, "// CalledWith:")
			} else if strings.HasPrefix(comment.Text, "// Example:") {
				example = strings.TrimPrefix(comment.Text, "// Example:")
			} else if strings.HasPrefix(comment.Text, "// ExpectedOutput:") {
				expectedOutput = strings.TrimPrefix(comment.Text, "// ExpectedOutput:")
			}
		}
		mdBuilder.WriteString(fmt.Sprintf("**Title:** %s\n\n**Description:** %s\n\n**Function:** %s\n\n**Called With:** %s\n\n**Example:** %s\n\n**Expected Output:** %s\n\n", title, description, function, calledWith, example, expectedOutput))
	}
}
