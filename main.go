func exportToCSV() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    outputFile, err := os.Create("export.csv")
    if err != nil {
        log.Println("Error creating CSV file:", err)
        return
    }
    defer outputFile.Close()

    writer := csv.NewWriter(outputFile)
    defer writer.Flush()

    scanner := bufio.NewScanner(inputFile)

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " ")
        if len(parts) > 1 {
            err := writer.Write(parts)
            if err != nil {
                log.Println("Error writing to CSV:", err)
                return
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }

    fmt.Println("History exported to export.csv.")
}
