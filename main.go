func sortLinesAlphabetically() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var lines []string
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    sort.Strings(lines)

    tempFile, err := os.Create("results_temp.txt")
    if err != nil {
        log.Println("Error creating temporary file:", err)
        return
    }
    defer tempFile.Close()

    for _, line := range lines {
        _, err := tempFile.WriteString(line + "\n")
        if err != nil {
            log.Println("Error writing to temporary file:", err)
        }
    }

    err = os.Rename("results_temp.txt", "results.txt")
    if err != nil {
        log.Println("Error replacing original file:", err)
    }

    fmt.Println("Sorted lines alphabetically successfully.")
}
