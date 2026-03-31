package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Stats struct {
	LinesRead    int
	LinesWritten int
	LinesRemoved int
	EmptyInput   bool
}

func upper(word string) string {
	return strings.ToUpper(word)
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

func isBlankOrDashes(line string) bool {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return true
	}

	for _, r := range trimmed {
		if r != '-' {
			return false
		}
	}
	return true
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

func addLineNumber(line string, number int) string {
	return fmt.Sprintf("%03d. %s", number, line)
}

func transformLine(line string) (string, bool) {
	line = strings.TrimSpace(line)

	if isBlankOrDashes(line) {
		return "", true
	}

	line = todoReplace(line)
	line = checkReverse(line)
	line = upper(line)

	return line, false
}

func buildOutput(processedLines []string, stats Stats) string {
	var result strings.Builder

	result.WriteString("SENTINEL FIELD REPORT — PROCESSED\n")
	result.WriteString("=================================\n")

	for _, line := range processedLines {
		result.WriteString(line)
		result.WriteString("\n")
	}

	result.WriteString("\n")
	result.WriteString("---------------\n")
	result.WriteString("FILE SUMMARY\n")
	result.WriteString("---------------\n")
	result.WriteString(fmt.Sprintf("Lines read    : %d\n", stats.LinesRead))
	result.WriteString(fmt.Sprintf("Lines written : %d\n", stats.LinesWritten))
	result.WriteString(fmt.Sprintf("Lines removed : %d\n", stats.LinesRemoved))
	result.WriteString("Rules applied : trim spaces, remove blank/dashes, replace TODO, reverse REVERSE lines, uppercase, line numbering\n")

	return result.String()
}

func transformation(inputFile string) ([]string, Stats, error) {
	var stats Stats
	var processedLines []string

	file, err := os.Open(inputFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, stats, fmt.Errorf("✗ File not found: %s", inputFile)
		}
		return nil, stats, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 1
	for scanner.Scan() {
		stats.LinesRead++
		line := scanner.Text()

		processed, removed := transformLine(line)
		if removed {
			stats.LinesRemoved++
			continue
		}

		numbered := addLineNumber(processed, lineNumber)
		processedLines = append(processedLines, numbered)

		stats.LinesWritten++
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return nil, stats, err
	}

	if stats.LinesRead == 0 {
		stats.EmptyInput = true
	}

	return processedLines, stats, nil
}

func applyRules(inputFile, outputFile string) (Stats, error) {
	var stats Stats

	same, err := sameFile(inputFile, outputFile)
	if err != nil {
		return stats, err
	}
	if same {
		return stats, fmt.Errorf("✗ Input and output cannot be the same file")
	}

	info, err := os.Stat(inputFile)
	if err != nil {
		if os.IsNotExist(err) {
			return stats, fmt.Errorf("✗ File not found: %s", inputFile)
		}
		return stats, err
	}
	if info.IsDir() {
		return stats, fmt.Errorf("✗ Input path is a directory, not a file: %s", inputFile)
	}

	outputInfo, err := os.Stat(outputFile)
	if err == nil && outputInfo.IsDir() {
		return stats, fmt.Errorf("✗ Cannot write to output: path is a directory, not a file")
	}

	processedLines, stats, err := transformation(inputFile)
	if err != nil {
		return stats, err
	}

	finalOutput := buildOutput(processedLines, stats)

	err = os.WriteFile(outputFile, []byte(finalOutput), 0644)
	if err != nil {
		return stats, fmt.Errorf("✗ Failed to write output file: %v", err)
	}

	return stats, nil
}

func printSummary(stats Stats) {
	fmt.Println("✦ Lines read    :", stats.LinesRead)
	fmt.Println("✦ Lines written :", stats.LinesWritten)
	fmt.Println("✦ Lines removed :", stats.LinesRemoved)
	fmt.Println("✦ Rules applied : trim spaces, remove blank/dashes, replace TODO, reverse REVERSE lines, uppercase, line numbering")

	if stats.EmptyInput {
		fmt.Println("⚠ Input file is empty. Nothing to process.")
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	stats, err := applyRules(inputFile, outputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printSummary(stats)
}