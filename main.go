package main

import (
	"fmt"

	"github.com/spf13/cobra"
)


var (
	uppercase bool
	lowercase bool
	titlecase bool
	wordcount bool
	charcount bool
)

var root_cmd = &cobra.Command {
	Use: "ritexfmt [text]",
	Short: "Minimal text formatter for fun.",
	Long: `ritexfmt is a CLI tool for formatting and analyzing text.
	You can provide text as arguments or pipe it in via stdin.
	Multiple formatting options can be applied simultaneously.`,
	Args: cobra.ArbitraryArgs,
	Run: run_text_formatter,
}
