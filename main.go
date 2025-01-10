func duplicateLines(lineNumber int) {
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

    scanner := bufio.NewScanner(inputFile)
    currentLine := 1
    for scanner.Scan() {
        line := scanner.Text()
        _, err := tempFile.WriteString(line + "\n")
        if err != nil {
            log.Println("Error writing to temporary file:", err)
        }

        if currentLine == lineNumber {
            _, err := tempFile.WriteString(line + "\n")
            if err != nil {
                log.Println("Error duplicating line in temporary file:", err)
            }
        }
        currentLine++
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    err = os.Rename("results_temp.txt", "results.txt")
    if err != nil {
        log.Println("Error replacing original file:", err)
    }

    fmt.Printf("Duplicated line %d in the file.\n", lineNumber)
}
