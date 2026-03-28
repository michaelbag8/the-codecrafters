THE BASE CONVERTER
  Concept: Number Systems & strconv

 Rules:
  → Write everything in Go. Standard library only.
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

  Requirements:

  → Support three input bases: hex, bin, dec.
  → For dec input, output both binary and hex.
  → For hex and bin input, output only decimal.
  → All hex output must be uppercase.
  → The program keeps running until: quit

  Validation — handle all of these cleanly:
  → "1G" is not valid hex.
  → "10201" is not valid binary.
  → "abc" is not a valid decimal.
  → Negative numbers must be supported for dec.
  → Empty input must not crash the program.

  Think about:
  → Which strconv functions handle base
    conversions for you?
  → How do you detect which characters are
    valid for a given base?
  → What is the difference between ParseInt
    and ParseUint — and does it matter here?

  ✦ Nail this and (hex) / (bin) in go-reloaded
    becomes a 5-minute task.
