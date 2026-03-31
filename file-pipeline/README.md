main()
 ├── validate arguments
 ├── check input/output paths
 ├── read input line by line
 ├── apply transformation pipeline
 ├── collect stats
 ├── write output file
 └── print terminal summary



validateArgs()
readLines()
applyRules()
writeOutput()
printSummary()



func main() {
    inputFile, outputFile, err := parseArgs(os.Args)
    if err != nil {
        fmt.Println(err)
        return
    }

    stats, processedLines, err := processFile(inputFile)
    if err != nil {
        fmt.Println(err)
        return
    }

    err = writeOutput(outputFile, processedLines)
    if err != nil {
        fmt.Println(err)
        return
    }

    printSummary(stats)
}


scanner := bufio.NewScanner(file)

for scanner.Scan() {
	line := scanner.Text()

	// apply rule 1
	// apply rule 2
	// apply rule 3
	// apply rule 4
	// apply rule 5

	// decide whether to keep or remove line
}