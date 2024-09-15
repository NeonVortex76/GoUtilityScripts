func countOperationsByType() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    operationCount := map[string]int{
        "+": 0,
        "-": 0,
        "*": 0,
        "/": 0,
        "%": 0,
        "^": 0,
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        for operation := range operationCount {
            if strings.Contains(line, " "+operation+" ") {
                operationCount[operation]++
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    fmt.Println("Operation counts:")
    for operation, count := range operationCount {
        fmt.Printf("%s: %d\n", operation, count)
    }
}
