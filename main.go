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
            inputFile.Close()
            tempFile.Close()
            return
        }

        writer.WriteString(line)
        if currentLine == lineNumber {
            writer.WriteString(line)
        }
        currentLine++
    }

    writer.Flush()
    if err1, err2 := inputFile.Close(), tempFile.Close(); err1 != nil || err2 != nil {
        log.Println("Error closing files")
        return
    }

    if err := os.Rename("results_temp.txt", "results.txt"); err != nil {
        log.Println("Error replacing original file:", err)
        return
    }

    log.Printf("Duplicated line %d in the file.\n", lineNumber)
}
