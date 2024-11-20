func filterLinesContainingWord() {
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

    var filterWord string
    fmt.Print("Enter the word to filter lines: ")
    fmt.Scanln(&filterWord)

    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, filterWord) {
            _, err := tempFile.WriteString(line + "\n")
            if err != nil {
                log.Println("Error writing to temporary file:", err)
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }

    err = os.Rename("results_temp.txt", "results.txt")
    if err != nil {
        log.Println("Error replacing original file:", err)
    }

    fmt.Println("Lines containing the specified word were filtered successfully.")
}
