func calculateAverageLineLength() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    scanner := bufio.NewScanner(inputFile)
    totalLength := 0
    lineCount := 0

    for scanner.Scan() {
        line := scanner.Text()
        totalLength += len(line)
        lineCount++
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }

    if lineCount > 0 {
        avgLength := float64(totalLength) / float64(lineCount)
        fmt.Printf("Average line length: %.2f characters\n", avgLength)
    } else {
        fmt.Println("No lines found in the file.")
    }
}
