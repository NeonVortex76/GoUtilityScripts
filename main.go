func findMostFrequentKeyword() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    keywordCount := make(map[string]int)
    scanner := bufio.NewScanner(inputFile)

    for scanner.Scan() {
        line := scanner.Text()
        words := strings.Fields(line)
        for _, word := range words {
            keywordCount[word]++
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    var mostFrequentKeyword string
    maxCount := 0
    for keyword, count := range keywordCount {
        if count > maxCount {
            mostFrequentKeyword = keyword
            maxCount = count
        }
    }

    fmt.Printf("Most frequent keyword: %s (Count: %d)\n", mostFrequentKeyword, maxCount)
}
