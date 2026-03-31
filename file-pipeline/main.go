package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func upper(word string) string {
	return strings.ToUpper(word)
}

func lower(word string) string {
	return strings.ToLower(word)
}

func todoReplace(word string)string{
	return strings.ReplaceAll(word, "TODO", "ACTION")
}

func reverse(word string)string{
	words := strings.Fields(word)
	for i, w := range words{
		runes := []rune(w)
		for left, right := 0, len(runes)-1; left < right; left, right = left + 1, right-1{
			runes[left], runes[right] = runes[right] , runes[left]
		}
		words[i] = string(runes)
	}
	return strings.Join(words," ")
}

func checkReserve(word string)string{
	words := strings.Fields(word)
	var result strings.Builder
	for _, w := range words {
		if w == "REVERSE"{
		change := reverse(w)
		result.WriteString(change)
	}
	 result.String()
	}
	return strings.Join(words, " ")
}

func applyRules(inputFile, outputFile string) error {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	text := string(content)
	result := transformation(text)

	return os.WriteFile(outputFile, []byte(result+"\n"), 0644)
}

func transformation(file string) string{

	files, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer files.Close()

	scanner := bufio.NewScanner(files)
	var result strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		caps := upper(line)
		cap := todoReplace(line)
		capz := checkReserve(line)
		result.WriteString(caps)
		result.WriteString(cap)
		result.WriteString(capz)
	}
	return result.String()

}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if err := applyRules(inputFile, outputFile); err != nil {
		fmt.Println("Error:", err)
	}
}
