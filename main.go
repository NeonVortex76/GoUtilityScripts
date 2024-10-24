func deleteOperationsBelowDurationThreshold() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var threshold float64
    fmt.Print("Enter the duration threshold: ")
    fmt.Scanln(&threshold)

    tempFile, err := os.Create("temp_results.txt")
    if err != nil {
        log.Println("Error creating temp file:", err)
        return
    }
    defer tempFile.Close()

    scanner := bufio.NewScanner(inputFile)
    writer := bufio.NewWriter(tempFile)

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " ")
        if len(parts) > 2 {
            duration, err := strconv.ParseFloat(parts[len(parts)-2], 64)
            if err != nil || duration >= threshold {
                _, err := writer.WriteString(line + "\n")
                if err != nil {
                    log.Println("Error writing to temp file:", err)
                }
            }
        }
    }

    err = writer.Flush()
    if err != nil {
        log.Println("Error flushing to temp file:", err)
    }

    os.Remove("results.txt")
    os.Rename("temp_results.txt", "results.txt")

    fmt.Println("Operations below the specified duration threshold have been deleted.")
}
