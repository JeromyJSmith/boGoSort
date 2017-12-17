package main

import (
    "fmt"
    "math/rand"
    "time"
    "os"
    "strconv"
)


func check_sorted(list []int) bool {
    for i := range list {
        if i == (len(list) - 1) {
            break
        }
        if list[i] > list[i+1] {
            return false
        }
    }
    return true
}


// Shuffle uses the Fisher-Yates Algorithm 
// https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
// to quickly create a randomized shuffle of the given list in O(n).
func shuffle(list []int) []int {
    last_index := len(list) - 1
    for i := last_index; i > 0; i-- {
        // Add 1 because Intn is non-inclusive on upper end
        j := rand.Intn(i + 1)
        tmp := list[j]
        list[j] = list[i]
        list[i] = tmp
    }
    return list
}

// trade time.Duration for float64
func bogosort(list []int) (int, float64) {
    permutations := 0

    t0 := time.Now()
    for ! check_sorted(list) {
        list = shuffle(list)
        permutations += 1
    }
    t1 := time.Now()
    duration := t1.Sub(t0)
    return permutations, duration.Seconds()
}


//Build_list returns a randomly shuffled list of numbers from 1 to len
func build_list(len int)[]int {
    list := make([]int, len)
    for i := 0; i < len; i++ {
        list[i] = i + 1
    }
    return shuffle(list)
}


// Implements Bogo sort, description from Wikipedia:
//  while not isInOrder(deck):
//      shuffle(deck)
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
        list := build_list(int(length))
        fmt.Printf("\nIteration: #%d\n", i+1)
        fmt.Println("Initial list:")
        fmt.Println(list)
        fmt.Println("------------------")
        results[i], duration = bogosort(list)
        total_permutations += results[i]
        total_time += duration
        fmt.Printf("Sorting finished in %d permutations\n", results[i])
        fmt.Printf("Duration: %fs\n", duration)

    }
    fmt.Printf("\n\nAverage over %d iterations: %.2f and %fs\n", iterations,
                                                        float64(total_permutations) / float64(iterations),
                                                        total_time / float64(iterations))
}

