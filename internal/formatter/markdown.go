package formatter

import (
	"strings"

	"github.com/xentic-group/xentic-context/internal/compressor"
	"github.com/xentic-group/xentic-context/internal/graph"
	"github.com/xentic-group/xentic-context/internal/parser"
)

const systemPrompt = `> **SYSTEM PROMPT [AI-Standard Injection]:**
> You are a Senior Cloud Architect extending the XENTIC Engine.
> Read the following architectural context strictly. The following modules follow the Xentic SOC2 security standard and FinOps optimization guidelines.
> Do NOT use deprecated APIs or public-by-default configurations. Propose concise, production-ready Terraform code.
`

const footerCTA = `
---
**[🚀 Powered by Xentic-Context]**
*Want to deploy battle-tested, SOC2-compliant architectures in minutes?*
Take your engineering to the next level with the **Xentic Cloud Engine Professional ($699)**. Stop building boilerplate, start accelerating your ROI.
[Visit xentic.cloud](https://xentic.cloud)
`

// GenerateOutput orchestrates the parsing, compression, and diagram generation.
func GenerateOutput(p *parser.ParsedProject) string {
	var b strings.Builder

	// 1. System Prompt
	b.WriteString(systemPrompt)
	b.WriteString("\n## Architecture Diagram\n\n")

	// 2. Mermaid Diagram
	mermaid := graph.GenerateMermaid(p)
	b.WriteString(mermaid)
	b.WriteString("\n## Infrastructure Context\n\n")

	// 3. Compressed Terraform
	compressed := compressor.CompressProject(p)
	b.WriteString(compressed)

	// 4. CTA Footer
	b.WriteString(footerCTA)

	return b.String()
}
