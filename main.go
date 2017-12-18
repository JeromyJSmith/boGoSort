package main

import (
    "fmt"
    "math/rand"
    "time"
    "os"
    "strconv"
    "github.com/I-Dont-Remember/bogosort/sort"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    var i int64
    if len(os.Args) < 3 {
        fmt.Println("Need [list length] [iterations]")
        return
    }
    length,_ := strconv.ParseInt(os.Args[1], 10, 64)
    iterations,_ := strconv.ParseInt(os.Args[2], 10, 64)
    if iterations < 1 || iterations < 1 {
        fmt.Println("Need plausible values")
    }

    fmt.Println("\n# # # # # # # #\nBegin boGo!\n------------------")
    results := make([]int, iterations)
    total_permutations := 0
    total_time := 0.0
    var duration float64 // time.Duration
    for i = 0; i < iterations; i++ {
        list := sort.Build_list(int(length))
        fmt.Printf("\nIteration: #%d\n", i+1)
        fmt.Println("Initial list:")
        fmt.Println(list)
        fmt.Println("------------------")
        results[i], duration = sort.Bogosort(list)
        total_permutations += results[i]
        total_time += duration
        fmt.Printf("Sorting finished in %d permutations\n", results[i])
        fmt.Printf("Duration: %fs\n", duration)

    }
    fmt.Printf("\n\nAverage over %d iterations: %.2f and %fs\n", iterations,
                                                        float64(total_permutations) / float64(iterations),
                                                        total_time / float64(iterations))
}

