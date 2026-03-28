package main

import (
	"fmt"
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
	for {
		var val string
		fmt.Print("Enter the value you want to convert: ")
		fmt.Scanln(&val)

		var base string
		fmt.Print("Enter the base (hex, dec, bin) or quit: ")
		fmt.Scanln(&base)

		if base == "" || base == "0" {
			fmt.Println("Base cannot be empty")
			continue
		}

		switch base {
		case "hex":
			num, err := toDecimal(val, 16)
			if err != nil {
				fmt.Println("Invalid hexadecimal number")
				continue
			}
			fmt.Println("Decimal:", num)

		case "bin":
			num, err := toDecimal(val, 2)
			if err != nil {
				fmt.Println("Invalid binary number")
				continue
			}
			fmt.Println("Decimal:", num)

		case "dec":
			num, err := strconv.ParseInt(val, 10, 64)
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
