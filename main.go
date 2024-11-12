func countLinesWithSpecificWord() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var word string
    fmt.Print("Enter the word to count its presence in lines: ")
    fmt.Scanln(&word)

    scanner := bufio.NewScanner(inputFile)
    lineCount := 0

    for scanner.Scan() {
        line := scanner.Text()
        words := strings.Fields(line)
        for _, w := range words {
            if w == word {
                lineCount++
                break
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }

    fmt.Printf("Number of lines containing the word '%s': %d\n", word, lineCount)
}
