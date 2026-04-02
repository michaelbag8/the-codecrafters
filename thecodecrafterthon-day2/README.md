THE BASE CONVERTER
  Concept: Number Systems & strconv

 How To Run The Program:
  There are two ways to run the base converter 
  1. Runing the .go file from it directory
  ```bash
      go run main.go
  ```
  There are 3 commands hex, bin and dec
   ```bash
      225 dec
   ```
   Output:
  ```bash
    Binary: 11100001
    Hex: E1
  ```
2. Runing the executable file ``convert``
```bash
    ./convert
```
 ```bash
      225 dec
   ```
   Output:
  ```bash
    Binary: 11100001
    Hex: E1
  ```
  → Must compile and run without errors.
  → Push to your the-codecrafters repo in a folder
    named: thecodecrafterthon-day2/
  → Include a README.md explaining how to run it

  go-reloaded needs to convert hex and binary
  strings into decimal numbers. This project
  teaches you exactly that — and makes you think
  about what happens when the input is garbage.

  Build a CLI tool that converts numbers between
  bases. It runs from the terminal like this:

     > convert 1E hex
       ✦ Decimal: 30

     > convert 10 bin
       ✦ Decimal: 2

     > convert 255 dec
       ✦ Binary:  11111111
       ✦ Hex:     FF

  
