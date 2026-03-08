# 🏰 Xentic-Context CLI
*"Stop Explaining. Start Architecting."*

Xentic-Context is a high-performance tool built in Go that acts as the ultimate bridge between your infrastructure code and Large Language Models (LLMs) such as Gemini, ChatGPT, and Claude.

Through a Structural Hydration process, the CLI transforms complex Terraform repositories into a single semantically optimized context document for AI, eliminating technical "noise" and prioritizing architectural logic.

## 🚀 Why Xentic-Context? (Viral Features)
- **📉 Token Compression:** Reduces context size by up to 60% by stripping unnecessary HCL metadata, enabling longer, deeper chats with AI.
- **🗺️ Auto-Mermaid Diagrams:** Automatically generates architecture flowcharts that AI can "visualize" to understand relationships between VPCs, Lambdas, and databases.
- **🛡️ AI-Standard Injection:** Automatically inserts System Prompts optimized with Xentic Group's AWS Security and FinOps standards.
- **⚡ Zero-Configuration:** Works instantly on any Terraform or OpenTofu project without prior setup.

## 🛠️ Quick Installation and Usage

### Prerequisites
- Go 1.21 or higher installed on your system.

### Installation
You can install `xentic-context` directly using `go install`:

```bash
go install github.com/xentic-group/xentic-context@latest
```
*(Make sure your `$(go env GOPATH)/bin` is in your system's `$PATH`)*

### Basic Usage
Navigate to the root of your Terraform or OpenTofu project and simply run:

```bash
xentic-context
```
This will instantly scan your project and generate an `ai-context.md` file. Then, simply drag and drop the `ai-context.md` file into your favorite AI chat (Claude, ChatGPT, Gemini) and start designing!

### Advanced Usage
You can customize the target directory and the output file name:
```bash
xentic-context -dir=/path/to/infrastructure -out=custom-context.md
```

## 🔒 Zero-Trust Security (Built-in)
You don't need to worry about accidentally leaking secrets to AI prompts. `xentic-context` includes a **Sensitive Data Filter** that proactively scans and redacts:
- Hardcoded AWS Access Keys and Secret Tokens.
- Attributes containing `password`, `secret`, `key`, or `token`.

All sensitive data is safely replaced with `"[REDACTED BY XENTIC-CONTEXT]"` before the file is generated. It also **strictly ignores** `.terraform/` directories and `.tfstate` files to prevent any remote state leakage.

## 💎 Elevate Your Engineering to the Next Level
Xentic-Context is the free gateway to the Xentic Group ecosystem. If you are a CTO or a Software Agency looking for bulletproof, audited, production-ready AWS foundations, you need the Xentic Cloud Engine:

- **Startup Edition ($349):** License for a single commercial project.
- **Professional Engine ($699):** Unlimited license for agencies. ROI focused on deployment speed.

## 🤝 Contributions and Community
This is an open-source project under the MIT license. We invite all cloud architects to contribute to improving the AI-Ready infrastructure standard.

### 🔗 Connect with Us
* **Website:** [xentic.cloud](https://www.xentic.cloud)
* **LinkedIn:** [Xentic Group](https://www.linkedin.com/company/xentic-group)
* **Contact:** [hello@xentic.cloud](mailto:hello@xentic.cloud)

*Developed with ❤️ by Xentic Group in Panama City 🇵🇦*
