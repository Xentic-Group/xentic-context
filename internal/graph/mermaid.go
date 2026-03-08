package graph

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/xentic-group/xentic-context/internal/parser"
)

// GenerateMermaid analyzes the project files to output a structural DAG for AI consumption.
func GenerateMermaid(p *parser.ParsedProject) string {
	var nodes []string
	var edges []string
    visitedEdges := make(map[string]bool)

	for _, file := range p.Files {
		for _, block := range file.Body.Blocks {
			if block.Type == "resource" && len(block.Labels) >= 2 {
				resType := block.Labels[0]
				resName := block.Labels[1]
				id := fmt.Sprintf("%s.%s", resType, resName)
				// Create standard node
				nodes = append(nodes, fmt.Sprintf("    %s[%s]", id, id))
				extractDependencies(block.Body, id, &edges, visitedEdges)
			} else if block.Type == "module" && len(block.Labels) >= 1 {
				modName := block.Labels[0]
				id := fmt.Sprintf("module.%s", modName)
				// Create rounded node for modules
				nodes = append(nodes, fmt.Sprintf("    %s((%s))", id, id))
				extractDependencies(block.Body, id, &edges, visitedEdges)
			}
		}
	}

	if len(nodes) == 0 {
		return "*(No resources found for auto-diagramming)*"
	}

	var b strings.Builder
	b.WriteString("```mermaid\ngraph TD\n")
	for _, n := range nodes {
		b.WriteString(n + "\n")
	}
	for _, e := range edges {
		b.WriteString(e + "\n")
	}
	b.WriteString("```\n")
	return b.String()
}

func extractDependencies(body *hclsyntax.Body, fromID string, edges *[]string, visited map[string]bool) {
	for _, attr := range body.Attributes {
		findRefsInExpr(attr.Expr, fromID, edges, visited)
	}
	for _, block := range body.Blocks {
		extractDependencies(block.Body, fromID, edges, visited)
	}
}

func findRefsInExpr(expr hclsyntax.Expression, fromID string, edges *[]string, visited map[string]bool) {
	for _, traverse := range expr.Variables() {
		if len(traverse) >= 2 {
			if rst, ok := traverse[0].(hcl.TraverseRoot); ok {
				// Ignore variables, locals, data sources for topological view clarity
				if rst.Name == "var" || rst.Name == "local" || rst.Name == "data" {
					continue
				}

				if attr, ok := traverse[1].(hcl.TraverseAttr); ok {
					targetID := fmt.Sprintf("%s.%s", rst.Name, attr.Name)
					link := fmt.Sprintf("    %s --> %s", fromID, targetID)
                    
                    if !visited[link] {
                        *edges = append(*edges, link)
                        visited[link] = true
                    }
				}
			}
		}
	}
}
