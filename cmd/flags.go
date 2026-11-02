package cmd

type options struct {
	upper, lower, title, reverse bool
	remove_spaces bool
	words, chars, palindrome bool
}


var opts options


func init() {
	flag := rootCmd.Flags();

	flag.BoolVarP(
		&opts.upper,
		"upper", "u",
		false,
		"Convert text to UPPERCASE",
	)
	flag.BoolVarP(
		&opts.lower,
		"lower", "l",
		false,
		"Convert text to lowercase",
	)
	flag.BoolVarP(
		&opts.title,
		"title", "t",
		false,
		"Convert text to Title Case",
	)
	flag.BoolVarP(
		&opts.reverse,
		"reverse", "r",
		false,
		"Reverse the text",
	)
	flag.BoolVarP(
		&opts.remove_spaces,
		"remove-spaces", "s",
		false,
		"Remove spaces and tabs",
	)

	flag.BoolVar(&opts.words, "words", false, "Count words")
	flag.BoolVar(&opts.chars, "chars", false, "Count characters")
	flag.BoolVar(&opts.palindrome, "palindrome", false, "Check if the text is a palindrome")

}
