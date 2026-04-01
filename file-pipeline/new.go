func checkUpperLine(line string) bool {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return false
	}

	hasLetter := false
	for _, r := range trimmed {
		if r >= 'a' && r <= 'z' {
			return false
		}
		if r >= 'A' && r <= 'Z' {
			hasLetter = true
		}
	}
	return hasLetter
}

func applyUpperCase(word string) string {
	if checkUpperLine(word) {
		return strings.ToLower(word)
	}
	return word
}

func checkReverse(line string) string {
	if strings.Contains(line, "REVERSE") {
		cleaned := strings.ReplaceAll(line, "REVERSE", "")
		cleaned = strings.Join(strings.Fields(cleaned), " ")
		return reverse(cleaned)
	}
	return line
}

func reverseWords(line string) string {
	words := strings.Fields(line)
	for left, right := 0, len(words)-1; left < right; left, right = left+1, right-1 {
		words[left], words[right] = words[right], words[left]
	}
	return strings.Join(words, " ")
}