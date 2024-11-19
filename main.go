func appendTimestampToLines() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    tempFile, err := os.Create("results_temp.txt")
    if err != nil {
        log.Println("Error creating temporary file:", err)
        return
    }
    defer tempFile.Close()

    currentTime := time.Now().Format("2006-01-02 15:04:05")

    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        line := scanner.Text()
        updatedLine := fmt.Sprintf("%s [%s]", line, currentTime)
        _, err := tempFile.WriteString(updatedLine + "\n")
        if err != nil {
            log.Println("Error writing to temporary file:", err)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }

    err = os.Rename("results_temp.txt", "results.txt")
    if err != nil {
        log.Println("Error replacing original file:", err)
    }

    fmt.Println("Timestamps appended to lines successfully.")
}
