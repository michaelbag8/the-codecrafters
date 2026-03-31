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

func todoReplace(word string) string {
	return strings.ReplaceAll(word, "TODO", "ACTION")
}

func reverseWords(line string) string {
	words := strings.Fields(line)
	for left, right := 0, len(words)-1; left < right; left, right = left+1, right-1 {
		words[left], words[right] = words[right], words[left]
	}
	return strings.Join(words, " ")
}

func checkReverse(line string) string {
	if strings.Contains(line, "REVERSE") {
		return reverseWords(line)
	}
	return line
}

func transformation(inputFile string) (string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		line = todoReplace(line)
		line = checkReverse(line)
		line = upper(line)

		result.WriteString(line)
		result.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return result.String(), nil
}

func applyRules(inputFile, outputFile string) error {
	result, err := transformation(inputFile)
	if err != nil {
		return err
	}

	return os.WriteFile(outputFile, []byte(result), 0644)
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