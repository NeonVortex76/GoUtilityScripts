func sortOperationsByDate() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var operations []string
    scanner := bufio.NewScanner(inputFile)

    for scanner.Scan() {
        operations = append(operations, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    sort.Slice(operations, func(i, j int) bool {
        layout := "2006-01-02" // Формат даты, например, 2024-10-03
        dateI := strings.Split(operations[i], " ")[0]
        dateJ := strings.Split(operations[j], " ")[0]
        timeI, errI := time.Parse(layout, dateI)
        timeJ, errJ := time.Parse(layout, dateJ)

        if errI != nil || errJ != nil {
            log.Println("Error parsing date:", errI, errJ)
            return false
        }
        return timeI.Before(timeJ)
    })

    tempFile, err := os.Create("sorted_results.txt")
    if err != nil {
        log.Println("Error creating temp file:", err)
        return
    }
    defer tempFile.Close()

    writer := bufio.NewWriter(tempFile)
    for _, operation := range operations {
        _, err := writer.WriteString(operation + "\n")
        if err != nil {
            log.Println("Error writing to temp file:", err)
        }
    }

    err = writer.Flush()
    if err != nil {
        log.Println("Error flushing to temp file:", err)
    }

    os.Remove("results.txt")
    os.Rename("sorted_results.txt", "results.txt")

    fmt.Println("Operations sorted by date.")
}
