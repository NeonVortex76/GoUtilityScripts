func reverseLineOrder() {
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

    tempFile, err := os.Create("results_temp.txt")
    if err != nil {
        log.Println("Error creating temporary file:", err)
        return
    }
    defer tempFile.Close()

    for i := len(lines) - 1; i >= 0; i-- {
        _, err := tempFile.WriteString(lines[i] + "\n")
        if err != nil {
            log.Println("Error writing to temporary file:", err)
        }
    }

    err = os.Rename("results_temp.txt", "results.txt")
    if err != nil {
        log.Println("Error replacing original file:", err)
    }

    fmt.Println("Reversed line order successfully.")
}
