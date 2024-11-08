func findKeywordFrequencyAcrossLines() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var keyword string
    fmt.Print("Enter the keyword to analyze frequency across lines: ")
    fmt.Scanln(&keyword)

    scanner := bufio.NewScanner(inputFile)
    lineCount := 0
    occurrenceCount := 0

    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, keyword) {
            occurrenceCount++
        }
        lineCount++
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }

    frequency := float64(occurrenceCount) / float64(lineCount) * 100
    fmt.Printf("Frequency of keyword '%s' across lines: %.2f%%\n", keyword, frequency)
}
