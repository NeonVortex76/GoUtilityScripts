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
    for currentLine := 1; scanner.Scan(); currentLine++ {
        line := scanner.Text()
        writeLineTwiceIfNeeded(tempFile, line, currentLine == lineNumber)
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    if err := os.Rename("results_temp.txt", "results.txt"); err != nil {
        log.Println("Error replacing original file:", err)
    }

    fmt.Printf("Duplicated line %d in the file.\n", lineNumber)
}

func writeLineTwiceIfNeeded(f *os.File, line string, duplicate bool) {
    f.WriteString(line + "\n")
    if duplicate {
        f.WriteString(line + "\n")
    }
}
