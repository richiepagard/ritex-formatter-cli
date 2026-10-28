# ritex-formatter-cli

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.24%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/Cobra-CLI-FF6B6B?style=for-the-badge&logo=go&logoColor=white" alt="Cobra" />
  <img src="https://img.shields.io/badge/Terminal-Tool-4CAF50?style=for-the-badge&logo=gnome-terminal&logoColor=white" alt="Terminal" />
  <img src="https://img.shields.io/badge/Format-Text-9C27B0?style=for-the-badge" alt="Format" />
</p>

Richie Text Formatter :| a minimal CLI for transforming and analyzing text from the terminal.

## Features

- Convert text to uppercase, lowercase, or title case
- Count words and characters
- Accept input from arguments or stdin
- Lightweight and fast with Go

## Usage

```bash
go run . "hello world" --upper
go run . --lower --words --chars

echo "Hello there" | go run . --title
```

## Options

- `-u, --upper` — convert to uppercase
- `-l, --lower` — convert to lowercase
- `-t, --title` — convert to title case
- `--words` — count words
- `--chars` — count characters

## Example

```bash
$ echo "hello from richie" | go run . --upper --words --chars
HELLO FROM RICHIE
Words: 4
Characters: 18
```

Made with a little Go magic and a lot of fun. 