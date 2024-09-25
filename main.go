func sortOperationsByDuration() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    type operationEntry struct {
        line     string
        duration time.Duration
    }

    var operations []operationEntry
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " ")
        if len(parts) < 3 {
            continue
        }

        durationStr := parts[2]
        duration, err := time.ParseDuration(durationStr)
        if err != nil {
            log.Println("Error parsing duration:", err)
            continue
        }

        operations = append(operations, operationEntry{line: line, duration: duration})
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    sort.Slice(operations, func(i, j int) bool {
        return operations[i].duration < operations[j].duration
    })

    fmt.Println("Operations sorted by duration:")
    for _, op := range operations {
        fmt.Println(op.line)
    }
}
