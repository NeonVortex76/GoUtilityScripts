func longestOperationLine() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    scanner := bufio.NewScanner(inputFile)
    longestLine := ""
    maxLength := 0

    for scanner.Scan() {
        line := scanner.Text()
        length := len(line)
        if length > maxLength {
            maxLength = length
            longestLine = line
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }

    fmt.Printf("Longest operation line (%d characters): %s\n", maxLength, longestLine)
}
