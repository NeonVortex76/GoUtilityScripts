func findLongestOperation() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    longestDuration := time.Duration(0)
    longestOperation := ""

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

        if duration > longestDuration {
            longestDuration = duration
            longestOperation = line
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    if longestOperation != "" {
        fmt.Printf("The longest operation is: %s (duration: %v)\n", longestOperation, longestDuration)
    } else {
        fmt.Println("No operations found.")
    }
}
