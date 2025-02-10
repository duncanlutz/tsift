# tsift

A command-line tool for generating Markdown documentation from TypeScript interface definitions. tsift scans TypeScript files and automatically creates formatted documentation for your interfaces, making it easier to maintain API documentation.

## Features

- Process single or multiple TypeScript files
- Scan entire directories recursively
- Generate formatted Markdown tables for interface properties
- Support for required and optional properties
- Custom output file support

## Installation

### Option 1: Pre-built Binaries

#### macOS
```bash
# Download the latest binary
curl -LO https://github.com/duncanlutz/tsift/releases/latest/download/tsift-darwin-amd64

# Make it executable
chmod +x tsift-darwin-amd64

# Move to a directory in your PATH (recommended location)
sudo mkdir -p /usr/local/bin
sudo mv tsift-darwin-amd64 /usr/local/bin/tsift

# Add to your shell configuration (~/.zshrc or ~/.bashrc)
echo 'export PATH="/usr/local/bin:$PATH"' >> ~/.zshrc  # for zsh
# OR
echo 'export PATH="/usr/local/bin:$PATH"' >> ~/.bashrc  # for bash

# Reload your shell configuration
source ~/.zshrc  # for zsh
# OR
source ~/.bashrc  # for bash
```

#### Linux
```bash
# Download the latest binary
curl -LO https://github.com/duncanlutz/tsift/releases/latest/download/tsift-linux-amd64

# Make it executable
chmod +x tsift-linux-amd64

# Move to a directory in your PATH (recommended location)
sudo mkdir -p /usr/local/bin
sudo mv tsift-linux-amd64 /usr/local/bin/tsift

# Add to your shell configuration
echo 'export PATH="/usr/local/bin:$PATH"' >> ~/.bashrc  # for bash
# OR
echo 'export PATH="/usr/local/bin:$PATH"' >> ~/.zshrc   # for zsh

# Reload your shell configuration
source ~/.bashrc  # for bash
# OR
source ~/.zshrc   # for zsh
```

### Option 2: Build from Source

If you prefer to build from source, you'll need Go installed on your system.

```bash
# Clone the repository
git clone https://github.com/duncanlutz/tsift.git
cd tsift

# Build the binary
go build -o tsift

# (Optional) Move to a directory in your PATH
sudo mv tsift /usr/local/bin/
```

## Usage

tsift can be run in two modes:

### Process specific files

```bash
tsift -f file1.ts,file2.ts,file3.ts
```

or

```bash
tsift --files file1.ts,file2.ts,file3.ts
```

### Process an entire directory

```bash
tsift -d ./src
```

or

```bash
tsift --directory ./src
```

### Specify output file

By default, tsift outputs to stdout. You can specify an output file using:

```bash
tsift -o output.md -f file1.ts
```

or

```bash
tsift --output output.md --directory ./src
```

## Output Format

The generated documentation follows this format:

```markdown
# Interface: InterfaceName

Interface description (if available)

| Property | Type | Required |
|----------|------|----------|
| propName | type | Yes/No   |
```

## Supported File Types

- `.ts` (TypeScript)
- `.tsx` (TypeScript React)

## Example

Given a TypeScript file with:

```typescript
interface User {
  id: number;
  name: string;
  email?: string;
}
```

tsift will generate:

```markdown
# Interface: User

| Property | Type | Required |
|----------|------|----------|
| id | number | Yes |
| name | string | Yes |
| email | string | No |
```

## Error Handling

- Exits with status code 1 if no input files or directory is specified
- Reports file processing errors and exits with status code 1
- Reports output file creation errors and exits with status code 1

## Contributing

Feel free to open issues and pull requests for additional features or bug fixes.

## License

MIT License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
