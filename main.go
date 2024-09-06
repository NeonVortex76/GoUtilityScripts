func operationStats() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    stats := map[string]int{
        "+": 0,
        "-": 0,
        "*": 0,
        "/": 0,
        "%": 0,
        "^": 0,
    }

    for scanner.Scan() {
        line := scanner.Text()
        for op := range stats {
            if strings.Contains(line, " "+op+" ") {
                stats[op]++
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    fmt.Println("Operation Statistics:")
    for op, count := range stats {
        fmt.Printf("%s: %d\n", op, count)
    }
}
