func editEntryByIndex() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var index int
    fmt.Print("Enter the index of the entry to edit: ")
    fmt.Scanln(&index)

    scanner := bufio.NewScanner(file)
    lines := []string{}
    lineNumber := 0

    for scanner.Scan() {
        line := scanner.Text()
        if lineNumber == index {
            fmt.Printf("Current entry at index %d: %s\n", index, line)
            fmt.Print("Enter new entry: ")
            var newEntry string
            fmt.Scanln(&newEntry)
            lines = append(lines, newEntry)
        } else {
            lines = append(lines, line)
        }
        lineNumber++
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    saveFixedEntries(lines)
    fmt.Printf("Entry at index %d has been updated.\n", index)
}
