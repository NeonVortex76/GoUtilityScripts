func duplicateLines(lineNumber int) {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }

    tempFile, err := os.Create("results_temp.txt")
    if err != nil {
        log.Println("Error creating temporary file:", err)
        inputFile.Close()
        return
    }

    scanner := bufio.NewScanner(inputFile)
    for currentLine := 1; scanner.Scan(); currentLine++ {
        line := scanner.Text()
        fmt.Fprintln(tempFile, line)
        if currentLine == lineNumber {
            fmt.Fprintln(tempFile, line)
        }
    }

    inputFile.Close()
    tempFile.Close()

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    if err := os.Rename("results_temp.txt", "results.txt"); err != nil {
        log.Println("Error replacing original file:", err)
    }

    fmt.Printf("Duplicated line %d in the file.\n", lineNumber)
}
