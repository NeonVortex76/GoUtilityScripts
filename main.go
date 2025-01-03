func countLinesContaining(substring string) int {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return 0
    }
    defer inputFile.Close()

    count := 0
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, substring) {
            count++
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return 0
    }

    fmt.Printf("The substring %q appears in %d lines.\n", substring, count)
    return count
}
