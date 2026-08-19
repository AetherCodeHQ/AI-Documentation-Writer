# AI Documentation Writer

Auto-generate comprehensive documentation for your codebases using AI - READMEs, API docs, and inline comments.

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)
[![CI](https://github.com/Qyroxen/ai-documentation-writer/actions/workflows/ci.yml/badge.svg)](https://github.com/Qyroxen/ai-documentation-writer/actions/workflows/ci.yml)

> Auto-generate comprehensive documentation for your codebases using AI - READMEs, API docs, and inline comments.

## What is it?

AI Documentation Writer is a command-line tool built with Go that helps developers auto-generate comprehensive documentation for your codebases using ai - readmes, api docs, and inline comments. It's designed to be fast, reliable, and easy to use.

## Why?

Every developer needs ai documentation writer — but existing tools are either too complex, too slow, or require cloud dependencies. We built AI Documentation Writer to be:
- **Fast** — Written in Go for maximum performance
- **Offline** — No cloud dependencies, your data stays on your machine
- **Simple** — Clean CLI interface with sensible defaults
- **Extensible** — Easy to customize and integrate into your workflow

## Features

- **Auto-generate README files** — Auto-generate README files
- **API documentation from code** — API documentation from code
- **Inline comment generation** — Inline comment generation
- **Multiple output formats (Markdown, HTML)** — Multiple output formats (Markdown, HTML)
- **Supports Go, Python, JavaScript, Java** — Supports Go, Python, JavaScript, Java
- **Customizable templates** — Customizable templates

## Quick Start

### Prerequisites

- Go 1.23 or later

### Install

```bash
# Install with go install
go install github.com/Qyroxen/ai-documentation-writer@latest

# Or build from source
git clone https://github.com/Qyroxen/ai-documentation-writer.git
cd ai-documentation-writer
go build -o ai-documentation-writer .
```

### Usage

```bash
# Basic usage
.ai-documentation-writer --help

# Example
./ai-documentation-writer generate --path ./src --format markdown
```

## Output

```
AI Documentation Writer v1.0.0

Scanning...

✓ Analysis complete
✓ Results ready

{
  "status": "success",
  "results": [...]
}
```

## Configuration

Create a `.config.yaml` file in your project root:

```yaml
# Configuration options
verbose: true
output: json
timeout: 30s
```

## CLI Flags

```
ai documentation writer [command]

Flags:
  --path string      Target path (default ".")
  --format string    Output format: json, text (default "text")
  --verbose          Enable verbose output
  --config string    Config file path
  --output string    Output file path
```

## Examples

### Basic Example

```bash
.ai-documentation-writer --path ./src
```

### Advanced Example

```bash
.ai-documentation-writer --path ./src --format json --output report.json --verbose
```

### CI/CD Integration

```yaml
# .github/workflows/ci.yml
- name: Run AI Documentation Writer
  run: |
    go install github.com/Qyroxen/ai-documentation-writer@latest
    ai-documentation-writer --path . --format json --output report.json
```

## Documentation

- [Getting Started](docs/getting-started.md)
- [Configuration](docs/configuration.md)
- [API Reference](docs/api-reference.md)
- [Examples](examples/)
- [Contributing](CONTRIBUTING.md)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Qyroxen** - [GitHub](https://github.com/Qyroxen)

---

**Found this useful?** Give it a ⭐ on GitHub!
