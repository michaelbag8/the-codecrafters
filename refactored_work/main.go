package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	if b == 0 {
		fmt.Printf("Not divisible by %d", b)
	}
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func calculator() {
	reader := bufio.NewReader(os.Stdin)
	for {
		//fmt.Println("Calculator Usage : add 3 5")
		fmt.Print(">> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Cannot be empty")
			continue
		}

		text := strings.Fields(input)
		command := strings.ToLower(text[0])

		switch command {
		case "add":
			num1, _ := strconv.Atoi(text[1])
			num2, _ := strconv.Atoi(text[2])
			fmt.Println("Result: ", add(num1, num2))
		case "sub":
			num1, _ := strconv.Atoi(text[1])
			num2, _ := strconv.Atoi(text[2])
			fmt.Println("Result: ", sub(num1, num2))
		case "mul":
			num1, _ := strconv.Atoi(text[1])
			num2, _ := strconv.Atoi(text[2])
			fmt.Println("Result: ", mult(num1, num2))
		}

	}
}

func main() {
	calculator()
}
