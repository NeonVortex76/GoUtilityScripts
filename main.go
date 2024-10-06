func importFromCSV() {
    inputFile, err := os.Open("import.csv")
    if err != nil {
        log.Println("Error opening CSV file:", err)
        return
    }
    defer inputFile.Close()

    outputFile, err := os.OpenFile("results.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Println("Error opening results file:", err)
        return
    }
    defer outputFile.Close()

    reader := csv.NewReader(inputFile)
    writer := bufio.NewWriter(outputFile)

    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Println("Error reading CSV:", err)
            return
        }

        line := strings.Join(record, " ") + "\n"
        _, err = writer.WriteString(line)
        if err != nil {
            log.Println("Error writing to results file:", err)
            return
        }
    }

    err = writer.Flush()
    if err != nil {
        log.Println("Error flushing data to file:", err)
    }

    fmt.Println("Operations imported from import.csv.")
}
