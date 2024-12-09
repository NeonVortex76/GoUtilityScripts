func prependLineNumbers() {
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

    scanner := bufio.NewScanner(inputFile)
    lineNumber := 1
    for scanner.Scan() {
        line := scanner.Text()
        numberedLine := fmt.Sprintf("%d: %s", lineNumber, line)
        _, err := tempFile.WriteString(numberedLine + "\n")
        if err != nil {
            log.Println("Error writing to temporary file:", err)
        }
        lineNumber++
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }

    err = os.Rename("results_temp.txt", "results.txt")
    if err != nil {
        log.Println("Error replacing original file:", err)
    }

    fmt.Println("Prepended line numbers successfully.")
}
