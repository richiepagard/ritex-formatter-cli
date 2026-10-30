package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"github.com/spf13/cobra"
)


var (
	uppercase bool
	lowercase bool
	titlecase bool
	reverse_text bool
	remove_space bool
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

func main() {
	if err := root_cmd.Execute(); err != nil {
		fmt.Println(err);

		// Execute exit code to exit the program
		os.Exit(1);
	}
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
	root_cmd.Flags().BoolVarP(
		&reverse_text,
		"reverse", "r",
		false,
		"Reversing the text",
	)
	root_cmd.Flags().BoolVarP(
		&remove_space,
		"remove-spaces", "s",
		false,
		"Removing spaces",
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

func run_text_formatter(cmd *cobra.Command, args []string) {
	var text string;

	// Check whether the input from function "args" or "stdin"
	if len(args) > 0 {
		text = strings.Join(args, " ");
	} else {
		// Read from "stdin"
		scanner := bufio.NewScanner(os.Stdin);
		var lines []string;

		for scanner.Scan() {
			lines = append(lines, scanner.Text());
		}
		text = strings.Join(lines, "\n");
	}

	if text == "" {
		fmt.Println("No input text provided");
		return
	}

	result := text;

	// Apply formatting transformations
	// Space remover only use if reverse text does not requested
	if uppercase {
		result = strings.ToUpper(result);
		fmt.Printf("Uppercase: %#v\n", result);
	}
	if lowercase {
		result = strings.ToLower(result);
		fmt.Printf("Lowercase: %#v\n", result);
	}
	if reverse_text {
		result = text_reverser(result);
		fmt.Printf("Reversed text: %#v\n", result);
	}
	if titlecase {
		result = strings.Title(result);
		fmt.Printf("Titlecase: %#v\n", result);
	}
	if remove_space && !reverse_text {
		result = space_remover(result);
		fmt.Printf("Removed spaces: %#v\n", result);
	}

	// Analysis if requested
	if wordcount {
		words := len(strings.Fields(text));
		fmt.Printf("Words: %d\n", words);
	}
	if charcount {
		chars := len(text);
		fmt.Printf("Characters: %d\n", chars);
	}
}


func text_reverser(text string) string {
	var (
		last int = len(text) - 1;
		first int = 0;
	)
	reversed := make([]byte, len(text));

	for char := last; char >= 0; char-- {
		reversed[first] = text[char];
		first++;
	}

	return string(reversed);
}

func space_remover(text string) (nospace string) {
	var replacer = strings.NewReplacer(" ", "", "\t", "");
	nospace = replacer.Replace(text);

	return
}
