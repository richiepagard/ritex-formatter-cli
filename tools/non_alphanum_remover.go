package tools

import (
	"regexp"
	"strings"
)


func NonAlphanumRemover(text string) (alphanum_text string) {
	/*
	Removing the non alphanumerical characters (e.g. ! : , ?) using regex.

	Arguments:
		text (string): The text which its non alphanum characters will removed.
	*/
	var non_alphanumeric_regex = regexp.MustCompile(`[^a-zA-Z0-9s]+`);

	alphanum_text = non_alphanumeric_regex.ReplaceAllString(text, "");
	alphanum_text = strings.ToLower(alphanum_text);

	return
}
