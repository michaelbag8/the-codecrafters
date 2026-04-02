package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLine(reader *bufio.Reader) string {
	lines, _ := reader.ReadString('\n')
	return strings.TrimSpace(lines)
}

func addition(a, b int) int {
	return a + b
}

func subtraction(a, b int) int {
	return a - b
}

func multiplication(a, b int) int {
	return a * b
}

func division(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("error: cannot divide by %v", b)
	}
	return a / b, nil
}

func help() {
	fmt.Println()
	fmt.Println("----help----")
	fmt.Println("Instructions")
	fmt.Println("1. Enter numbers")
	fmt.Println("2. Select an operation")
	fmt.Println("3. Enter help for instructions")
	fmt.Println("4. Enter quit to exit")
	fmt.Println()
}

func calculator() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command (e.g. add 4 5): ")
		input := readLine(reader)

		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		switch command {
		case "quit":
			fmt.Println("Thank you for using my calculator, goodbye!!!")
			return

		case "help":
			help()
			continue
		}

		if len(parts) != 3 {
			fmt.Println("Invalid format. Use: add 4 5")
			continue
		}

		a, err1 := strconv.Atoi(parts[1])
		b, err2 := strconv.Atoi(parts[2])

		if err1 != nil || err2 != nil {
			fmt.Println("Please enter valid numbers.")
			continue
		}

		switch command {
		case "add":
			fmt.Println("Result:", addition(a, b))

		case "sub":
			fmt.Println("Result:", subtraction(a, b))

		case "mul":
			fmt.Println("Result:", multiplication(a, b))

		case "div":
			result, err := division(a, b)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Result:", result)
			}

		default:
			fmt.Println("Unknown operation. Use add, sub, mul, div, help, quit.")
		}
	}
}

func main() {
	calculator()
}
