func filterHistoryByOperation(operation string) {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    filtered := []string{}

    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, " "+operation+" ") {
            filtered = append(filtered, line)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    if len(filtered) == 0 {
        fmt.Printf("No operations found for '%s'.\n", operation)
    } else {
        fmt.Printf("Filtered history for operation '%s':\n", operation)
        for _, line := range filtered {
            fmt.Println(line)
        }
    }
}
