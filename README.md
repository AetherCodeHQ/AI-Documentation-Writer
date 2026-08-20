# AI Documentation Writer

![CI](https://github.com/Qyroxen/AI-Documentation-Writer/actions/workflows/ci.yml/badge.svg) ![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go) ![License](https://img.shields.io/badge/License-MIT-yellow.svg) ![Stars](https://img.shields.io/github/stars/Qyroxen/AI-Documentation-Writer?style=social)

> Auto-generate documentation for your codebase using AI

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/AI-Documentation-Writer?style=social)](https://github.com/Qyroxen/AI-Documentation-Writer/stargazers)

## What is it?

AI Documentation Writer analyzes your code and generates comprehensive documentation including API docs, README files, and inline comments.

## Why should you care?

Writing documentation is tedious. Let AI do it while you focus on coding.

## Demo

```bash
./ai-doc-writer generate --path ./my-project
```

**Output:**
```
Generated documentation:
  - README.md (updated)
  - docs/api-reference.md
  - 45 inline comments added
```

## Features

- Auto-generate README files
- Create API reference documentation
- Add inline code comments
- Support for 10+ languages
- Customizable documentation style

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/AI-Documentation-Writer.git
cd AI-Documentation-Writer
go build -o ai-doc-writer .

# Run
./ai-doc-writer --path ./my-project
```

## CLI Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--path` | Source code directory | `.` |
| `--output` | Output directory | `./docs` |
| `--style` | Documentation style (concise, detailed) | `concise` |
| `--lang` | Target language | `en` |

## Examples

# Generate docs
./ai-doc-writer generate --path ./src

# Detailed style
./ai-doc-writer generate --path ./src --style detailed

# Spanish docs
./ai-doc-writer generate --path ./src --lang es

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/AI-Documentation-Writer/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/AI-Documentation-Writer?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/AI-Documentation-Writer/network/members">
    <img src="https://img.shields.io/github/forks/Qyroxen/AI-Documentation-Writer?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/AI-Documentation-Writer/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/AI-Documentation-Writer" alt="Issues">
  </a>
</p>
