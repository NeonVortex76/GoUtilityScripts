func sortHistory(order string) {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lines := []string{}

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    sort.SliceStable(lines, func(i, j int) bool {
        timeI := getTimeFromEntry(lines[i])
        timeJ := getTimeFromEntry(lines[j])

        if order == "asc" {
            return timeI.Before(timeJ)
        }
        return timeI.After(timeJ)
    })

    fmt.Println("Sorted History:")
    for _, line := range lines {
        fmt.Println(line)
    }
}

func getTimeFromEntry(entry string) time.Time {
    parts := strings.Split(entry, " - ")
    if len(parts) < 1 {
        return time.Time{}
    }

    timestamp := parts[0]
    parsedTime, err := time.Parse(time.RFC3339, timestamp)
    if err != nil {
        log.Println("Error parsing time:", err)
        return time.Time{}
    }

    return parsedTime
}
