package main

import (
	"fmt"
	"os"
)

func readFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	return string(data)
}

func writeFile(filename string, content string) {
	err := os.WriteFile(filename, []byte(content+"\n"), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
	}

	input1 := os.Args[1]
	input2 := os.Args[2]

	content := readFile(input1)
	content = applyTransformation(content)
	writeFile(input2, content)
}
