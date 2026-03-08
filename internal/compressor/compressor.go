package compressor

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/xentic-group/xentic-context/internal/parser"
)

var (
	// awsKeyRegex detects AKIA/ASIA keys strictly
	awsKeyRegex = regexp.MustCompile(`(?i)(AKIA|ASIA)[A-Z0-9]{16}`)
)

// CompressProject takes a ParsedProject and returns a completely compressed
// single string representing the infrastructure, stripped of generic noise.
func CompressProject(p *parser.ParsedProject) string {
	var buf bytes.Buffer

	for path, file := range p.Files {
		buf.WriteString(fmt.Sprintf("\n### File: %s\n```hcl\n", path))
		buf.WriteString(compressBody(file.Body, file.Source, 0))
		buf.WriteString("```\n")
	}

	return buf.String()
}

func compressBody(body *hclsyntax.Body, source []byte, depth int) string {
	var buf bytes.Buffer
	for _, block := range body.Blocks {
		writeBlock(&buf, block, source, depth)
	}
	return buf.String()
}

func writeBlock(buf *bytes.Buffer, block *hclsyntax.Block, source []byte, depth int) {
	indent := strings.Repeat("  ", depth)
	buf.WriteString(indent + block.Type + " ")
	
	for _, label := range block.Labels {
		buf.WriteString(fmt.Sprintf("%q ", label))
	}
	buf.WriteString("{\n")

	// Write attributes compactly
	for name, attr := range block.Body.Attributes {
		if isSensitiveName(name) {
			buf.WriteString(fmt.Sprintf("%s  %s = \"[REDACTED BY XENTIC-CONTEXT]\"\n", indent, name))
		} else {
			// Extract raw value without comments
			rawExpr := extractRaw(attr.Expr.Range(), source)
			// Proactively redact any AWS key inside the expression
			redactedExpr := awsKeyRegex.ReplaceAllString(rawExpr, "\"[REDACTED BY XENTIC-CONTEXT]\"")
			buf.WriteString(fmt.Sprintf("%s  %s = %s\n", indent, name, redactedExpr))
		}
	}

	// Recursively write nested blocks
	for _, child := range block.Body.Blocks {
		writeBlock(buf, child, source, depth+1)
	}

	buf.WriteString(indent + "}\n")
}

func isSensitiveName(name string) bool {
	n := strings.ToLower(name)
	return strings.Contains(n, "password") || strings.Contains(n, "secret") || 
	       strings.Contains(n, "token") || strings.Contains(n, "key")
}

func extractRaw(r hcl.Range, source []byte) string {
	if r.Start.Byte < 0 || r.End.Byte > len(source) || r.Start.Byte >= r.End.Byte {
		return `""` // fallback
	}
	val := string(source[r.Start.Byte:r.End.Byte])
	return val
}
