func viewOperationsByDurationRange() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var minDuration, maxDuration float64
    fmt.Print("Enter the minimum duration: ")
    fmt.Scanln(&minDuration)
    fmt.Print("Enter the maximum duration: ")
    fmt.Scanln(&maxDuration)

    scanner := bufio.NewScanner(inputFile)

    fmt.Printf("Operations with duration between %.2f and %.2f:\n", minDuration, maxDuration)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " ")
        if len(parts) > 2 {
            duration, err := strconv.ParseFloat(parts[len(parts)-2], 64)
            if err == nil && duration >= minDuration && duration <= maxDuration {
                fmt.Println(line)
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }
}
