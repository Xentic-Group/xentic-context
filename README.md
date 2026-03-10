# ⚡ context-ai
*"Stop Explaining. Start Architecting."*

**context-ai** is a high-performance CLI tool built in Go — the ultimate bridge between your infrastructure code and Large Language Models (LLMs) such as Gemini, ChatGPT, and Claude.

Through a **Structural Hydration** process, the CLI transforms complex Terraform repositories into a single semantically optimized context document for AI, eliminating technical "noise" and prioritizing architectural logic.

---

## 🚀 Features

- **📉 Token Compression:** Reduces context size by up to 60% by stripping unnecessary HCL metadata, enabling longer, deeper chats with AI.
- **🗺️ Auto-Mermaid Diagrams:** Automatically generates architecture flowcharts that AI can "visualize" to understand relationships between VPCs, Lambdas, and databases.
- **🛡️ AI-Standard Injection:** Automatically inserts System Prompts optimized with Xentic Group's AWS Security and FinOps standards.
- **⚡ Zero-Configuration:** Works instantly on any Terraform or OpenTofu project without prior setup.
- **🔒 Zero-Trust Security:** Sensitive data (passwords, tokens, keys) is automatically redacted before being written to the output file.

---

## 🛠️ Installation

### ✅ Recommended: Homebrew (macOS & Linux)
The easiest and most universal way to install `context-ai` — no Go required.

```bash
brew install Xentic-Group/tap/context-ai
```

### Alternative: Go Install
If you have Go 1.21+ installed on your system:

```bash
go install github.com/xentic-group/xentic-context@latest
```
*(Make sure your `$(go env GOPATH)/bin` is in your system's `$PATH`)*

---

## 🚀 Usage

Navigate to the root of your Terraform or OpenTofu project and run:

```bash
context-ai
```

This will scan your project and generate an `ai-context.md` file. Simply drag and drop it into your favorite AI chat (Claude, ChatGPT, Gemini) and start designing.

### Flags

| Flag | Description | Default |
|------|-------------|---------|
| `-dir` | Directory containing Terraform/OpenTofu files | `.` (current dir) |
| `-out` | Output markdown file name | `ai-context.md` |
| `-h / --help` | Show help and usage examples | — |

### Examples

```bash
# Scan the current directory (most common):
context-ai

# Scan a specific infrastructure folder:
context-ai -dir=/path/to/infrastructure

# Custom output name:
context-ai -dir=./infra -out=context-for-claude.md
```

---

## 🔒 Zero-Trust Security (Built-in)

`context-ai` includes a **Sensitive Data Filter** that proactively scans and redacts:
- Hardcoded AWS Access Keys and Secret Tokens.
- Attributes containing `password`, `secret`, `key`, or `token`.

All sensitive data is safely replaced with `"[REDACTED BY XENTIC-CONTEXT]"` before the file is generated. It also **strictly ignores** `.terraform/` directories and `.tfstate` files to prevent any remote state leakage.

---

## 💎 Take Your Engineering to the Next Level

**context-ai** is the free gateway to the Xentic Group ecosystem. If you are a CTO or a Software Agency looking for bulletproof, audited, production-ready AWS foundations:

- **🟢 Startup Edition ($349):** License for a single commercial project. Save 200+ hours of AWS architecture.
- **🔵 Professional Engine ($699):** Unlimited license for agencies. ROI focused on deployment speed.

👉 [xentic.cloud](https://www.xentic.cloud)

---

## 🤝 Contributing

This is an open-source project under the MIT license. We invite all cloud architects to contribute to improving the AI-Ready infrastructure standard.

## 🔗 Connect with Us
* **Website:** [xentic.cloud](https://www.xentic.cloud)
* **LinkedIn:** [Xentic Group](https://www.linkedin.com/company/xentic-group)
* **Contact:** [hello@xentic.cloud](mailto:hello@xentic.cloud)

*Developed with ❤️ by Xentic Group in Panama City 🇵🇦*
