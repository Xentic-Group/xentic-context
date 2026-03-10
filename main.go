package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/xentic-group/xentic-context/internal/formatter"
	"github.com/xentic-group/xentic-context/internal/parser"
)

const banner = `
╔═══════════════════════════════════════════════════════════════╗
║              ⚡ context-ai by Xentic Group ⚡                 ║
║     Stop Explaining. Start Architecting.                      ║
╚═══════════════════════════════════════════════════════════════╝

  Transforms your Terraform / OpenTofu repository into a single
  AI-optimized context document for Claude, ChatGPT, and Gemini.

  USAGE:
    context-ai [flags]

  FLAGS:
    -dir  string    Path to Terraform project (default: current dir)
    -out  string    Output file name (default: ai-context.md)

  EXAMPLES:
    # Scan the current directory:
    context-ai

    # Scan a specific folder:
    context-ai -dir=/path/to/infrastructure

    # Custom output filename:
    context-ai -dir=./infra -out=context-for-claude.md

  FEATURES:
    ✅  Token compression (up to 60%% less noise for LLMs)
    ✅  Auto Mermaid architecture diagrams
    ✅  Sensitive data redaction (Zero‑Trust)
    ✅  AI System‑Prompt injection (Xentic SOC2 standard)

  MORE TOOLS:
    🔗  https://xentic.cloud
    📧  hello@xentic.cloud
`

func main() {
	var dir string
	var out string
	var help bool

	flag.StringVar(&dir, "dir", ".", "Directory containing Terraform/OpenTofu files")
	flag.StringVar(&out, "out", "ai-context.md", "Output markdown file name")
	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&help, "h", false, "Show help (shorthand)")

	// Override the default usage function with the branded banner
	flag.Usage = func() {
		fmt.Println(banner)
	}

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	start := time.Now()

	fmt.Println("╔═══════════════════════════════════╗")
	fmt.Println("║  ⚡ context-ai  by Xentic Group   ║")
	fmt.Println("╚═══════════════════════════════════╝")
	fmt.Printf("\n🔍 Scanning directory: %s\n", dir)

	project, err := parser.ParseDirectory(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error parsing directory: %v\n", err)
		os.Exit(1)
	}

	if len(project.Files) == 0 {
		fmt.Println("⚠️  No Terraform (.tf) files found in the specified directory.")
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
	fmt.Println("\n💡 Tip: Drag & drop the generated file into Claude, ChatGPT or Gemini to start architecting.")
	fmt.Println("🔗 xentic.cloud")

	// Performance Benchmark Check
	if duration.Seconds() > 2.0 {
		fmt.Printf("⚠️  Warning: Execution time exceeded the 2.0s baseline (took %v).\n", duration)
	}
}
