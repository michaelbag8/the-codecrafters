package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toDecimal(str string, base int) (int64, error) {
	convert, err := strconv.ParseInt(str, base, 64)
	return convert, err
}

func decimalToHexBin(num int64, base int) string {
	convert := strconv.FormatInt(num, base)
	return strings.ToUpper(convert)
}

func baseConverter() {
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("Base Converter - Online. The Usage: 255 dec")
		fmt.Print(">> ")

		val, _ := reader.ReadString('\n')
		val = strings.TrimSpace(val)

		if val == "" {
			fmt.Println("Invalid input. Use format: <value> <base>")
			continue
		}

		txt := strings.Fields(val)

		if len(txt) < 2 {
			fmt.Println("Invalid input. Use format: <value> <base>")
			continue
		}

		number := txt[0]
		convert := strings.ToLower(txt[1])

		if convert == "" || convert == "0" {
			fmt.Println("Base cannot be empty")
			continue
		}

		switch convert {
		case "hex":
			num, err := toDecimal(number, 16)
			if err != nil {
				fmt.Println("Invalid hexadecimal number")
				continue
			}
			fmt.Println("Decimal:", num)

		case "bin":
			num, err := toDecimal(number, 2)
			if err != nil {
				fmt.Println("Invalid binary number")
				continue
			}
			fmt.Println("Decimal:", num)

		case "dec":
			num, err := strconv.ParseInt(number, 10, 64)
			if err != nil {
				fmt.Println("Invalid decimal number")
				continue
			}
			fmt.Println("Binary:", decimalToHexBin(num, 2))
			fmt.Println("Hex:", decimalToHexBin(num, 16))

		case "quit":
			fmt.Println("Thank you for testing the base converter")
			return

		default:
			fmt.Println("Invalid base. Use hex, bin, dec, or quit")
		}
	}
}

func main() {
	baseConverter()
}
