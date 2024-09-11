func backupHistory() {
    srcFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening source file:", err)
        return
    }
    defer srcFile.Close()

    backupFile, err := os.Create("results_backup.txt")
    if err != nil {
        log.Println("Error creating backup file:", err)
        return
    }
    defer backupFile.Close()

    scanner := bufio.NewScanner(srcFile)
    for scanner.Scan() {
        line := scanner.Text()
        if _, err := backupFile.WriteString(line + "\n"); err != nil {
            log.Println("Error writing to backup file:", err)
            return
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading source file:", err)
        return
    }

    fmt.Println("Backup created successfully.")
}
