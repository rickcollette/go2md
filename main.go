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

	if node.Doc != nil {
		mdBuilder.WriteString("# " + node.Name.Name + "\n\n")
		for _, comment := range node.Doc.List {
			mdBuilder.WriteString(comment.Text + "\n")
		}
		mdBuilder.WriteString("\n")
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

func processFuncDecl(decl *ast.FuncDecl, mdBuilder *strings.Builder) {
	mdBuilder.WriteString(fmt.Sprintf("## Function: %s\n\n", decl.Name.Name))
	if decl.Doc != nil {
		var title, description, function, calledWith, expectedOutput string
		for _, comment := range decl.Doc.List {
			if strings.HasPrefix(comment.Text, "// Title:") {
				title = strings.TrimPrefix(comment.Text, "// Title:")
			} else if strings.HasPrefix(comment.Text, "// Description:") {
				description = strings.TrimPrefix(comment.Text, "// Description:")
			} else if strings.HasPrefix(comment.Text, "// Function:") {
				function = strings.TrimPrefix(comment.Text, "// Function:")
			} else if strings.HasPrefix(comment.Text, "// CalledWith:") {
				calledWith = strings.TrimPrefix(comment.Text, "// CalledWith:")
			} else if strings.HasPrefix(comment.Text, "// ExpectedOutput:") {
				expectedOutput = strings.TrimPrefix(comment.Text, "// ExpectedOutput:")
			}
		}
		mdBuilder.WriteString(fmt.Sprintf("**Title:** %s\n\n**Description:** %s\n\n**Function:** %s\n\n**Called With:** %s\n\n**Expected Output:** %s\n\n", title, description, function, calledWith, expectedOutput))
	}
}