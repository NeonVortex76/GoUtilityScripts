func calculateAverageDuration() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var totalDuration float64
    var count int
    scanner := bufio.NewScanner(inputFile)

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " ")
        if len(parts) > 2 {
            duration, err := strconv.ParseFloat(parts[len(parts)-2], 64)
            if err == nil {
                totalDuration += duration
                count++
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }

    if count > 0 {
        avgDuration := totalDuration / float64(count)
        fmt.Printf("Average operation duration: %.2f\n", avgDuration)
    } else {
        fmt.Println("No valid operations found for calculating average duration.")
    }
}
