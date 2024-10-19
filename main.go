func countOperationsByDateRange() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var startDateStr, endDateStr string
    fmt.Print("Enter the start date (YYYY-MM-DD): ")
    fmt.Scanln(&startDateStr)
    fmt.Print("Enter the end date (YYYY-MM-DD): ")
    fmt.Scanln(&endDateStr)

    layout := "2006-01-02"
    startDate, err := time.Parse(layout, startDateStr)
    if err != nil {
        log.Println("Invalid start date:", err)
        return
    }
    endDate, err := time.Parse(layout, endDateStr)
    if err != nil {
        log.Println("Invalid end date:", err)
        return
    }

    scanner := bufio.NewScanner(inputFile)
    count := 0

    for scanner.Scan() {
        line := scanner.Text()
        dateStr := strings.Split(line, " ")[0]
        operationDate, err := time.Parse(layout, dateStr)
        if err != nil {
            log.Println("Error parsing date:", err)
            continue
        }

        if (operationDate.After(startDate) && operationDate.Before(endDate)) || operationDate.Equal(startDate) || operationDate.Equal(endDate) {
            count++
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }

    fmt.Printf("Total operations between %s and %s: %d\n", startDateStr, endDateStr, count)
}
