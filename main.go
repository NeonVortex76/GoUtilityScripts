func countWordOccurrences() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    wordCounts := make(map[string]int)
    scanner := bufio.NewScanner(inputFile)

    for scanner.Scan() {
        line := scanner.Text()
        words := strings.Fields(line)
        for _, word := range words {
            wordCounts[word]++
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }

    fmt.Println("Word occurrences:")
    for word, count := range wordCounts {
        fmt.Printf("%s: %d\n", word, count)
    }
}
