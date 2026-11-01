package analyze

import "strings"


func Words(text string) int {
	return len(strings.Fields(text));
}

func Chars(text string) int {
	return len(text);
}
