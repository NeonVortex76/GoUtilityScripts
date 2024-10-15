func filterOperationsByKeyword() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var keyword string
    fmt.Print("Enter the keyword to filter by: ")
    fmt.Scanln(&keyword)

    scanner := bufio.NewScanner(inputFile)
    fmt.Printf("Operations containing '%s':\n", keyword)

    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, keyword) {
            fmt.Println(line)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }
}
