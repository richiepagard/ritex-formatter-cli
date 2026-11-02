package cmd

import (
	"fmt"
	"os"
	"io"
	"strings"
	"errors"
)
import "github.com/spf13/cobra"
import (
	"ritexfmt/internal/format"
	"ritexfmt/internal/analyze"
)


func run(cmd *cobra.Command, args []string) error {
	text, err := read_input(args);

	if err != nil {
		return err;
	}
	if text == "" {
		return errors.New("No input text provided");
	}

	result := text;

	// Apply formatting transformations
	// Space remover only use if reverse text does not requested
	if opts.upper {
		result = format.Upper(result);
		fmt.Printf("Uppercase: %#v\n", result);
	}
	if opts.lower {
		result = format.Lower(result);
		fmt.Printf("Lowercase: %#v\n", result);
	}
	if opts.reverse {
		result = format.TextReverse(result);
		fmt.Printf("Reversed text: %#v\n", result);
	}
	if opts.title {
		result = format.Title(result);
		fmt.Printf("Titlecase: %#v\n", result);
	}
	if opts.remove_spaces && !opts.reverse {
		result = format.SpaceRemover(result);
		fmt.Printf("Removed spaces: %#v\n", result);
	}

	// Analysis if requested
	if opts.words {
		words := analyze.Words(result);
		fmt.Printf("Words: %d\n", words);
	}
	if opts.chars {
		chars := analyze.Chars(result);
		fmt.Printf("Chars: %d\n", chars);
	}
	if opts.palindrome {
		palindrome_status, _ := analyze.ValidPalindrome(result);

		switch palindrome_status {
			case false:
				fmt.Println("Not a palindrome text");
			case true:
				fmt.Println("Valid palindrome text");
		}
	}

	return nil
}

func read_input(args []string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil;
	}

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", err;
	}

	return strings.TrimRight(string(data), "\n"), nil;
}
