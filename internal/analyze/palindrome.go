package analyze

import "ritexfmt/tools"


func ValidPalindrome(text string) (bool, string) {
	palindrome_text := text;
	text = tools.NonAlphanumRemover(text);
	var first, last int = 0, len(text)-1;

	for first < last {
		if text[first] == text[last] {
			first++;
			last--;
		} else {
			break;
		}
	}

	if first >= last {
		return true, palindrome_text;
	}

	return false, palindrome_text;
}
