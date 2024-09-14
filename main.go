func fixInvalidEntries() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lines := []string{}
    lineNumber := 0
    hasErrors := false

    for scanner.Scan() {
        line := scanner.Text()
        lineNumber++

        if !isValidEntryFormat(line) {
            fmt.Printf("Invalid format at line %d: %s\n", lineNumber, line)
            fmt.Print("Please enter a corrected entry: ")
            var correctedEntry string
            fmt.Scanln(&correctedEntry)
            lines = append(lines, correctedEntry)
            hasErrors = true
        } else {
            lines = append(lines, line)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    if hasErrors {
        saveFixedEntries(lines)
        fmt.Println("Invalid entries were fixed.")
    } else {
        fmt.Println("No invalid entries found.")
    }
}

func saveFixedEntries(lines []string) {
    tempFile, err := os.Create("results_temp.txt")
    if err != nil {
        log.Println("Error creating temporary file:", err)
        return
    }
    defer tempFile.Close()

    for _, line := range lines {
        if _, err := tempFile.WriteString(line + "\n"); err != nil {
            log.Println("Error writing to temporary file:", err)
            return
        }
    }

    err = os.Remove("results.txt")
    if err != nil {
        log.Println("Error deleting original file:", err)
        return
    }

    err = os.Rename("results_temp.txt", "results.txt")
    if err != nil {
        log.Println("Error renaming temporary file:", err)
        return
    }
}
