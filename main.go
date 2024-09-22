func calculateMedianResult() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var results []float64
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " ")
        if len(parts) < 3 {
            continue
        }

        result, err := strconv.ParseFloat(parts[2], 64) // Предполагаем, что результат операции — третья часть строки
        if err != nil {
            log.Println("Error parsing result:", err)
            continue
        }

        results = append(results, result)
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    if len(results) == 0 {
        fmt.Println("No results found.")
        return
    }

    sort.Float64s(results)

    var median float64
    n := len(results)
    if n%2 == 0 {
        median = (results[n/2-1] + results[n/2]) / 2
    } else {
        median = results[n/2]
    }

    fmt.Printf("The median result is: %.2f\n", median)
}
