func operationDurationStats() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    operationDurations := make(map[string]time.Duration)
    operationCounts := make(map[string]int)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " ")
        if len(parts) < 3 {
            continue
        }

        operation := parts[1]
        durationStr := parts[2]
        duration, err := time.ParseDuration(durationStr)
        if err != nil {
            log.Println("Error parsing duration:", err)
            continue
        }

        operationDurations[operation] += duration
        operationCounts[operation]++
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    fmt.Println("Operation duration statistics:")
    for operation, totalDuration := range operationDurations {
        count := operationCounts[operation]
        avgDuration := totalDuration / time.Duration(count)
        fmt.Printf("%s: total time = %v, average time = %v (%d occurrences)\n", operation, totalDuration, avgDuration, count)
    }
}
