package main

import (
    "fmt"
    "math/rand"
    "time"
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


func bogosort(list []int) {
    fmt.Println("Begin boGo!\n------------------")
    fmt.Println("Initial list:")
    fmt.Println(list)
    fmt.Println("------------------")
    permutations := 0

    for ! check_sorted(list) {
        list = shuffle(list)
        permutations += 1
    }
    fmt.Printf("Sorting finished in %d permutations\n", permutations)
}


// Implements Bogo sort, description from Wikipedia:
//  while not isInOrder(deck):
//      shuffle(deck)
func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    list := []int{2, 5, 1, 4, 3}
    bogosort(list)
}

