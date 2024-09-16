func mostFrequentOperation() {
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

    mostFrequent := ""
    maxCount := 0
    for operation, count := range operationCount {
        if count > maxCount {
            mostFrequent = operation
            maxCount = count
        }
    }

    if mostFrequent != "" {
        fmt.Printf("The most frequently used operation is '%s' with %d occurrences.\n", mostFrequent, maxCount)
    } else {
        fmt.Println("No operations found.")
    }
}
