package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func readLine(reader *bufio.Reader) string{
	lines, _:= reader.ReadString('\n')
	return strings.TrimSpace(lines)
}

func readInteger(reader *bufio.Reader, prompt string) int{
	for {
		fmt.Print(prompt)
		input := readLine(reader)
		number, err := strconv.Atoi(input)
		if err != nil{
			fmt.Println("Please enter a number: ")
			continue
		}
		return number
	}
}

func addition(a,b int)int{
	return a + b
}

func subtraction(a,b int)int{
	return a - b
}

func multiplication(a,b int)int{
	return a * b
}

func division(a,b int) (int, error){
	if b == 0{
		return 0, fmt.Errorf("error: cannot divide by %v", b)
	}
	return a / b, nil
}

func help(){
	fmt.Println()
	fmt.Println("----help----")
	fmt.Println("Instructions")
	fmt.Println("1. Enter numbers")
	fmt.Println("2. Select an operation")
	fmt.Println("3. Enter help for instructions")
	fmt.Println("4. Enter quit to exit")
	fmt.Println()
}
func calculator(){
	reader := bufio.NewReader(os.Stdin)
	for {
		a := readInteger(reader, "Enter the first number: ")
		b := readInteger(reader, "Enter the second number: ")

		fmt.Print("Select an operation (add, sub, mul, div, help, quit): ")
		operation := readLine(reader)

		switch operation {
		case "add":
			fmt.Println("Result: ", addition(a,b))

		case "sub":
			fmt.Println("Result: ", subtraction(a,b))

		case "mul":
			fmt.Println("Result: ", multiplication(a,b))

		case "div":
			result, err := division(a,b)
			if err != nil{
				fmt.Printf("cannot divide by %v: \n", b)
			}else{
				fmt.Println("Result: ", result)
			}
		case "help":
			help()
			continue

		case "quit":
			fmt.Println("Thank you for using my calculator, goodbye!!!")
			return
		default:
			fmt.Printf("%q is an invalid operation: acceptable operations are (add,mul,sub,div,help,quit)", operation)	
		}
		fmt.Println("We go again...")
	}
}

func main(){
	calculator()
}