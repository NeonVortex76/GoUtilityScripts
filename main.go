func findLongestLine() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var longestLine string
    maxLen := 0

    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) > maxLen {
            longestLine = line
            maxLen = len(line)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    fmt.Printf("Longest line (%d characters): %s\n", maxLen, longestLine)
}
