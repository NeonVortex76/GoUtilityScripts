func filterOperationsByResultThreshold() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var threshold float64
    fmt.Print("Enter the result threshold to filter operations: ")
    fmt.Scanln(&threshold)

    scanner := bufio.NewScanner(file)
    fmt.Printf("Operations with results greater than %.2f:\n", threshold)

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " ")
        if len(parts) < 2 {
            continue
        }

        result, err := strconv.ParseFloat(parts[1], 64)
        if err != nil {
            log.Println("Error
