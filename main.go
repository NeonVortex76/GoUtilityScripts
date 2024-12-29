func countWordOccurrences(word string) int {
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
        count += strings.Count(line, word)
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return 0
    }

    fmt.Printf("The word '%s' appears %d times in the file.\n", word, count)
    return count
}
