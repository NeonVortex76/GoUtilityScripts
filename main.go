func filterOperationsByKeywordCount() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var keyword string
    var minCount, maxCount int
    fmt.Print("Enter the keyword to filter by: ")
    fmt.Scanln(&keyword)
    fmt.Print("Enter minimum count: ")
    fmt.Scanln(&minCount)
    fmt.Print("Enter maximum count: ")
    fmt.Scanln(&maxCount)

    operationCounts := make(map[string]int)
    scanner := bufio.NewScanner(inputFile)

    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, keyword) {
            operationCounts[line]++
        }
    }

    for operation, count := range operationCounts {
        if count >= minCount && count <= maxCount {
            fmt.Printf("%s - Count: %d\n", operation, count)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }
}
