func findLinesWithMultipleKeywords() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var keyword1, keyword2 string
    fmt.Print("Enter the first keyword: ")
    fmt.Scanln(&keyword1)
    fmt.Print("Enter the second keyword: ")
    fmt.Scanln(&keyword2)

    scanner := bufio.NewScanner(inputFile)

    fmt.Printf("Lines containing both '%s' and '%s':\n", keyword1, keyword2)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, keyword1) && strings.Contains(line, keyword2) {
            fmt.Println(line)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }
}
