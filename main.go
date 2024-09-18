func countOperationsByDateRange() {
    layout := "2006-01-02"
    var startDateStr, endDateStr string

    fmt.Print("Enter start date (YYYY-MM-DD): ")
    fmt.Scanln(&startDateStr)
    startDate, err := time.Parse(layout, startDateStr)
    if err != nil {
        fmt.Println("Invalid start date format.")
        return
    }

    fmt.Print("Enter end date (YYYY-MM-DD): ")
    fmt.Scanln(&endDateStr)
    endDate, err := time.Parse(layout, endDateStr)
    if err != nil {
        fmt.Println("Invalid end date format.")
        return
    }

    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    operationCount := 0

    for scanner.Scan() {
        line := scanner.Text()
        dateStr := strings.Split(line, " ")[0] // Предполагаем, что дата стоит первой в строке
        operationDate, err := time.Parse(layout, dateStr)
        if err != nil {
            continue
        }

        if operationDate.After(startDate) && operationDate.Before(endDate) || operationDate.Equal(startDate) || operationDate.Equal(endDate) {
            operationCount++
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }

    fmt.Printf("Total operations between %s and %s: %d\n", startDateStr, endDateStr, operationCount)
}
