package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"path/filepath"
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


func sameFile(inputFile, outputFile string) (bool, error) {
	absInput, err := filepath.Abs(inputFile)
	if err != nil {
		return false, err
	}

	absOutput, err := filepath.Abs(outputFile)
	if err != nil {
		return false, err
	}

	return absInput == absOutput, nil
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
