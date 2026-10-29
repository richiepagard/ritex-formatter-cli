# ritex-formatter-cli
> Richie Text Formatter CLI

Richie Text Formatter :|

a minimal CLI for transforming and analyzing text from the terminal.

[![Golang](https://go-skill-icons.vercel.app/api/icons?i=go)](https://go.dev/doc/) __
[![Environment](https://go-skill-icons.vercel.app/api/icons?i=bash,powershell)](https://cobra.dev/)

## Features

- Convert text to uppercase, lowercase, or title case
- Count words and characters
- Accept input from arguments or `stdin`
- Lightweight and fast with **Go**

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
