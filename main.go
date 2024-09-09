func averageResult() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var total float64
    var count int

    for scanner.Scan() {
        line := scanner.Text()
        result, err := extractResultFromEntry(line)
        if err == nil {
            total += result
            count++
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    if count == 0 {
        fmt.Println("No numeric results found.")
        return
    }

    average := total / float64(count)
    fmt.Printf("The average result of all operations is: %.2f\n", average)
}

func extractResultFromEntry(entry string) (float64, error) {
    parts := strings.Split(entry, " = ")
    if len(parts) < 2 {
        return 0, fmt.Errorf("invalid entry format")
    }

    result, err := strconv.ParseFloat(parts[1], 64)
    if err != nil {
        return 0, fmt.Errorf("invalid result format")
    }

    return result, nil
}
