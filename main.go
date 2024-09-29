func deleteOperationsByDurationRange() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var minDurationStr, maxDurationStr string
    fmt.Print("Enter minimum duration to delete (e.g., 2s): ")
    fmt.Scanln(&minDurationStr)
    fmt.Print("Enter maximum duration to delete (e.g., 5s): ")
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

    var updatedLines []string
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " ")
        if len(parts) < 3 {
            updatedLines = append(updatedLines, line)
            continue
        }

        durationStr := parts[2]
        duration, err := time.ParseDuration(durationStr)
        if err != nil {
            log.Println("Error parsing duration:", err)
            updatedLines = append(updatedLines, line)
            continue
        }

        if duration < minDuration || duration > maxDuration {
            updatedLines = append(updatedLines, line)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    file.Close()

    newFile, err := os.Create("results.txt")
    if err != nil {
        log.Println("Error creating file:", err)
        return
    }
    defer newFile.Close()

    writer := bufio.NewWriter(newFile)
    for _, line := range updatedLines {
        _, err := writer.WriteString(line + "\n")
        if err != nil {
            log.Println("Error writing to file:", err)
            return
        }
    }

    if err := writer.Flush(); err != nil {
        log.Println("Error flushing data to file:", err)
    }

    fmt.Println("Operations within the specified duration range have been deleted.")
}
