func listOperationsExceedingThreshold() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var threshold float64
    fmt.Print("Enter duration threshold: ")
    fmt.Scanln(&threshold)

    scanner := bufio.NewScanner(inputFile)

    fmt.Printf("Operations exceeding duration threshold of %.2f:\n", threshold)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " ")
        if len(parts) > 2 {
            duration, err := strconv.ParseFloat(parts[len(parts)-2], 64)
            if err == nil && duration > threshold {
                fmt.Println(line)
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }
}
