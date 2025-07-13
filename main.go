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

    reader := bufio.NewReader(inputFile)
    writer := bufio.NewWriter(tempFile)

    currentLine := 1
    for {
        line, err := reader.ReadString('\n')
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Println("Error reading line:", err)
            return
        }

        writer.WriteString(line)
        if currentLine == lineNumber {
            writer.WriteString(line)
        }
        currentLine++
    }

    if err := writer.Flush(); err != nil {
        log.Println("Error flushing to file:", err)
        return
    }

    if err := os.Rename("results_temp.txt", "results.txt"); err != nil {
        log.Println("Error replacing original file:", err)
        return
    }

    log.Printf("Duplicated line %d in the file.\n", lineNumber)
}
