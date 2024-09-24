func findShortestOperation() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    shortestDuration := time.Duration(math.MaxInt64)
    shortestOperation := ""

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

        if duration < shortestDuration {
            shortestDuration = duration
            shortestOperation = line
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    if shortestOperation != "" {
        fmt.Printf("The shortest operation is: %s (duration: %v)\n", shortestOperation, shortestDuration)
    } else {
        fmt.Println("No operations found.")
    }
}
