func filterOperationsByDurationRange() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var minDurationStr, maxDurationStr string
    fmt.Print("Enter minimum duration (e.g., 2s): ")
    fmt.Scanln(&minDurationStr)
    fmt.Print("Enter maximum duration (e.g., 5s): ")
    fmt.Scanln(&maxDurationStr)

    minDuration, err := time.ParseDuration(minDurationStr)
    if err != nil {
        log.Println("Error parsing minimum duration:", err)
        return
    }

    maxDuration, err := time.ParseDuration(maxDurationStr)
    if err != nil {
        log.Println("Error parsing maximum duration:", err)
        return
    }

    scanner := bufio.NewScanner(file)
    fmt.Printf("Operations with duration between %v and %v:\n", minDuration, maxDuration)

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

        if duration >= minDuration && duration <= maxDuration {
            fmt.Println(line)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }
}
