func countOperationsContainingKeyword() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var keyword string
    fmt.Print("Enter the keyword to search for: ")
    fmt.Scanln(&keyword)

    scanner := bufio.NewScanner(inputFile)
    count := 0

    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, keyword) {
            count++
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }

    fmt.Printf("Number of operations containing '%s': %d\n", keyword, count)
}
