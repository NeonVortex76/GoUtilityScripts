func validateHistory() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    isValid := true
    lineNumber := 0

    for scanner.Scan() {
        line := scanner.Text()
        lineNumber++

        if !isValidEntryFormat(line) {
            fmt.Printf("Invalid format at line %d: %s\n", lineNumber, line)
            isValid = false
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    if isValid {
        fmt.Println("All entries are valid.")
    } else {
        fmt.Println("Some entries have invalid format.")
    }
}

func isValidEntryFormat(entry string) bool {
    parts := strings.Split(entry, " = ")
    if len(parts) != 2 {
        return false
    }

    _, err := strconv.ParseFloat(parts[1], 64)
    if err != nil {
        return false
    }

    return true
}
