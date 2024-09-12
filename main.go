func restoreHistoryFromBackup() {
    backupFile, err := os.Open("results_backup.txt")
    if err != nil {
        log.Println("Error opening backup file:", err)
        return
    }
    defer backupFile.Close()

    restoreFile, err := os.Create("results.txt")
    if err != nil {
        log.Println("Error creating results file:", err)
        return
    }
    defer restoreFile.Close()

    scanner := bufio.NewScanner(backupFile)
    for scanner.Scan() {
        line := scanner.Text()
        if _, err := restoreFile.WriteString(line + "\n"); err != nil {
            log.Println("Error writing to results file:", err)
            return
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading backup file:", err)
        return
    }

    fmt.Println("History restored successfully from backup.")
}
