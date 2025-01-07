func appendTimestampToFile() {
    file, err := os.OpenFile("results.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    timestamp := time.Now().Format("2006-01-02 15:04:05")
    _, err = file.WriteString("Timestamp: " + timestamp + "\n")
    if err != nil {
        log.Println("Error writing to file:", err)
    }

    fmt.Println("Appended current timestamp to file.")
}
