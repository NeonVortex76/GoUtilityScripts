func deleteEntryByIndex() {
    fmt.Print("Enter the index of the entry to delete (starting from 0): ")
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

    deletedEntry := lines[index]

    lines = append(lines[:index], lines[index+1:]...)

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

    // Сохраняем удалённую запись в отдельный файл
    saveDeletedEntry(deletedEntry)

    fmt.Println("Entry deleted successfully.")
}

func saveDeletedEntry(entry string) {
    file, err := os.OpenFile("deleted_entries.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Println("Error opening file for deleted entries:", err)
        return
    }
    defer file.Close()

    if _, err := file.WriteString(entry + "\n"); err != nil {
        log.Println("Error saving deleted entry:", err)
    }
}

func restoreDeletedEntry() {
    fmt.Print("Enter the index of the deleted entry to restore (starting from 0): ")
    var index int
    fmt.Scanln(&index)

    file, err := os.Open("deleted_entries.txt")
    if err != nil {
        log.Println("Error opening deleted entries file:", err)
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

    restoredEntry := lines[index]

    resultFile, err := os.OpenFile("results.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Println("Error opening results file:", err)
        return
    }
    defer resultFile.Close()

    if _, err := resultFile.WriteString(restoredEntry + "\n"); err != nil {
        log.Println("Error writing restored entry:", err)
        return
    }

    fmt.Println("Entry restored successfully.")
}
