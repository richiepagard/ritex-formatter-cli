package tools

import "strings"


func SpaceRemover(text string) (nospace string) {
	/*
	Removing the spaces and tabs in a text provided by function args.

	Arguments:
		text (string): The input text will remove its spaces and tabs.
	*/
	var replacer = strings.NewReplacer(" ", "", "\t", "");
	nospace = replacer.Replace(text);

	return
}
