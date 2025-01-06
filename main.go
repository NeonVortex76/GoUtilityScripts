func getWordCount() int {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return 0
    }
    defer inputFile.Close()

    wordCount := 0
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        words := strings.Fields(scanner.Text())
        wordCount += len(words)
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return 0
    }

    fmt.Printf("The file contains %d words.\n", wordCount)
    return wordCount
}
