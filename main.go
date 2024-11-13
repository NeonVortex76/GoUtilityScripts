func duplicateLineChecker() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    scanner := bufio.NewScanner(inputFile)
    lineSet := make(map[string]bool)
    duplicates := make([]string, 0)

    for scanner.Scan() {
        line := scanner.Text()
        if lineSet[line] {
            duplicates = append(duplicates, line)
        } else {
            lineSet[line] = true
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }

    if len(duplicates) > 0 {
        fmt.Println("Duplicate lines found:")
        for _, dup := range duplicates {
            fmt.Println(dup)
        }
    } else {
        fmt.Println("No duplicate lines found.")
    }
}
