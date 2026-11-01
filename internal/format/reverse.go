package format


func TextReverse(text string) string {
	/*
	Saves the reversed version of the text into a new string.
	Returns the reversed string.

	Arguments:
		text (string): The text will saves its reverse version in result.
	*/
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
