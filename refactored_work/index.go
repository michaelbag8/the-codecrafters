package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Upper(word string) string {
	return strings.ToUpper(word)
}

func Lower(word string) string {
	return strings.ToLower(word)
}

func Cap(word string) string {
	words := strings.Fields(word)

	for i, r := range words {
		runes := []rune(strings.ToLower(r))
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

func Reverse(word string) string {
	words := strings.Fields(word)

	for i, wo := range words {
		runes := []rune(wo)
		for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
			runes[l], runes[r] = runes[r], runes[l]
		}
		words[i] = string(runes)

	}
	return strings.Join(words, " ")
}

func snake(word string) string {
	var result strings.Builder
	words := []rune(word)

	for _, ch := range words {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) || unicode.IsSpace(ch) || ch == '_' {
			if unicode.IsLetter(ch) {
				result.WriteRune(unicode.ToLower(ch))
			} else {
				result.WriteRune(ch)
			}

		}

	}

	word = result.String()
	word = strings.Join(strings.Fields(word), "_")
	return word

}

var small_words = []string{"a", "an", "the", "and", "but", "or", "for", "nor", "on", "at", "to", "by", "in", "of", "up", "as", "is", "it"}

func checksmallwordz(word string) bool {
	for _, w := range small_words {
		if w == word {
			return true
		}
	}
	return false
}

func titlez(word string) string {
	words := strings.Fields(word)

	for i, r := range words {
		lower := strings.ToLower(r)
		if i == 0 || !checksmallwordz(lower) {
			runes := []rune(lower)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		} else {
			words[i] = lower
		}
	}
	return strings.Join(words, " ")
}

func transformer() {
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("SENTINEL IS ONLINE: Usage upper text")

		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "" {
			fmt.Println("cannot be empty, please enter a command and text")
			continue
		}

		textInputs := strings.Fields(text)
		commands := strings.ToLower(textInputs[0])

		if commands == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			return
		}

		if len(textInputs) < 2 {
			fmt.Println("No text provided ")
			continue
		}

		texts := strings.Join(textInputs[1:], " ")

		switch commands {
		case "upper":
			fmt.Println("¬", Upper(texts))
		case "lower":
			fmt.Println("¬", Lower(texts))
		case "snake":
			fmt.Println("¬", snake(texts))
		case "title":
			fmt.Println("¬", titlez(texts))
		case "cap":
			fmt.Println("¬", Cap(texts))
		default:
			fmt.Println("Unknown command, please enter acceptable command ")
		}

	}
}

func main() {
	transformer()
}
