func replaceWordInLines() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    tempFile, err := os.Create("results_temp.txt")
    if err != nil {
        log.Println("Error creating temporary file:", err)
        return
    }
    defer tempFile.Close()

    var oldWord, newWord string
    fmt.Print("Enter the word to replace: ")
    fmt.Scanln(&oldWord)
    fmt.Print("Enter the new word: ")
    fmt.Scanln(&newWord)

    scanner := bufio.NewScanner(inputFile)

    for scanner.Scan() {
        line := scanner.Text()
        updatedLine := strings.ReplaceAll(line, oldWord, newWord)
        _, err := tempFile.WriteString(updatedLine + "\n")
        if err != nil {
            log.Println("Error writing to temporary file:", err)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }

    err = os.Rename("results_temp.txt", "results.txt")
    if err != nil {
        log.Println("Error replacing original file:", err)
    }

    fmt.Println("Word replacement completed successfully.")
}
