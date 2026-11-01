package format

import "strings"


func Upper(text string) string {
	return strings.ToUpper(text);
}

func Lower(text string) string {
	return strings.ToLower(text);
}

func Title(text string) string {
	return strings.Title(text);
}
