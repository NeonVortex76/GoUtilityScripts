func editEntryByIndex() {
    fmt.Print("Enter the index of the entry to edit (starting from 0): ")
    var index int
    fmt.Scanln(&index)

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

    if index < 0 || index >= len(lines) {
        fmt.Println("Invalid index.")
        return
    }

    fmt.Printf("Current entry: %s\n", lines[index])
    fmt.Print("Enter the new entry: ")
    var newEntry string
    fmt.Scanln(&newEntry)

    lines[index] = newEntry

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

    fmt.Println("Entry edited successfully.")
}
