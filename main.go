func countOperationsByKeyword() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var keyword string
    fmt.Print("Enter the keyword to count operations: ")
    fmt.Scanln(&keyword)

    scanner := bufio.NewScanner(file)
    count := 0

    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, keyword) {
            count++
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    fmt.Printf("Total operations containing the keyword '%s': %d\n", keyword, count)
}
