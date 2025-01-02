func findLongestLine() string {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return ""
    }
    defer inputFile.Close()

    var longestLine string
    maxLength := 0

    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) > maxLength {
            maxLength = len(line)
            longestLine = line
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return ""
    }

    fmt.Printf("The longest line is: %q\n", longestLine)
    return longestLine
}
