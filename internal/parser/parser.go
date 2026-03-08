package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// ParsedFile holds both the AST and the raw source for slicing
type ParsedFile struct {
	Body   *hclsyntax.Body
	Source []byte
}

// ParsedProject holds the AST of all parsed Terraform files.
type ParsedProject struct {
	Files map[string]*ParsedFile
}

// shouldIgnore encapsulates the "Ignorer Logic" to prevent token bloat
func shouldIgnore(path string) bool {
	// Exclude restricted directories
	if strings.Contains(path, string(filepath.Separator)+".terraform"+string(filepath.Separator)) ||
		strings.Contains(path, string(filepath.Separator)+".git"+string(filepath.Separator)) {
		return true
	}
	
	base := filepath.Base(path)
	// Strictly ignore lock files and state files
	if base == ".terraform.lock.hcl" || strings.HasSuffix(base, ".tfstate") || strings.HasSuffix(base, ".tfstate.backup") {
		return true
	}
	
	// Only process .tf files (and optionally .tfvars)
	if !strings.HasSuffix(base, ".tf") && !strings.HasSuffix(base, ".tfvars") {
		return true
	}
	
	return false
}

// ParseDirectory recursively scans and parses HCL files
func ParseDirectory(dir string) (*ParsedProject, error) {
	parser := hclparse.NewParser()
	project := &ParsedProject{
		Files: make(map[string]*ParsedFile),
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden folders proactively
		if info.IsDir() {
			if info.Name() == ".terraform" || info.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}

		if shouldIgnore(path) {
			return nil
		}

		file, diags := parser.ParseHCLFile(path)
		if diags.HasErrors() {
			fmt.Fprintf(os.Stderr, "Warning: failed to parse %s: %s\n", path, diags.Error())
			return nil // proceed with what we extracted
		}

		if body, ok := file.Body.(*hclsyntax.Body); ok {
			// Store relative path for cleaner output
			relPath, _ := filepath.Rel(dir, path)
			project.Files[relPath] = &ParsedFile{
				Body:   body,
				Source: file.Bytes,
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return project, nil
}
