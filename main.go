func filterOperationsByDate() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var dateInput string
    fmt.Print("Enter the date (YYYY-MM-DD) to filter operations: ")
    fmt.Scanln(&dateInput)

    scanner := bufio.NewScanner(file)
    fmt.Printf("Operations on %s:\n", dateInput)

    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, dateInput) {
            fmt.Println(line)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }
}
