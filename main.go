package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/xentic-group/xentic-context/internal/formatter"
	"github.com/xentic-group/xentic-context/internal/parser"
)

func main() {
	var dir string
	var out string
	flag.StringVar(&dir, "dir", ".", "Directory containing Terraform/OpenTofu files")
	flag.StringVar(&out, "out", "ai-context.md", "Output markdown file name")
	flag.Parse()

	start := time.Now()

	fmt.Printf("🔍 Scanning directory: %s\n", dir)
	project, err := parser.ParseDirectory(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error parsing directory: %v\n", err)
		os.Exit(1)
	}

	if len(project.Files) == 0 {
		fmt.Println("⚠️ No Terraform (.tf) files found in the specified directory.")
		os.Exit(0)
	}

	fmt.Printf("🧩 Structural Hydration in progress (%d files)...\n", len(project.Files))
	markdown := formatter.GenerateOutput(project)

	err = os.WriteFile(out, []byte(markdown), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error writing output file: %v\n", err)
		os.Exit(1)
	}

	duration := time.Since(start)
	
	fmt.Printf("✅ Successfully created '%s' in %v!\n", out, duration)
	
	// Performance Benchmark Check
	if duration.Seconds() > 2.0 {
		fmt.Printf("⚠️ Warning: Execution time exceeded the 2.0s baseline (took %v).\n", duration)
	}
}
