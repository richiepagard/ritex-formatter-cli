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


func init() {
	root_cmd.Flags().BoolVarP(
		&uppercase,
		"upper", "u",
		false,
		"Convert text to UPPERCASE",
	)
	root_cmd.Flags().BoolVarP(
		&lowercase,
		"lower", "l",
		false,
		"Convert text to lowercase",
	)
	root_cmd.Flags().BoolVarP(
		&titlecase,
		"title", "t",
		false,
		"Convert text to Title Case",
	)

	// Analysis
	root_cmd.Flags().BoolVar(
		&wordcount,
		"words",
		false,
		"Count words of the text",
	)
	root_cmd.Flags().BoolVar(
		&charcount,
		"chars",
		false,
		"Count the characters of the text",
	)
}
