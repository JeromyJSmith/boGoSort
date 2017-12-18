package sort

import (
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
// https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_Shuffle
// to quickly create a randomized Shuffle of the given list in O(n).
func Shuffle(list []int) []int {
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

// Bogosort implements this description from Wikipedia:
//  while not isInOrder(deck):
//      Shuffle(deck)
// returning permutations tried and time taken in seconds.
func Bogosort(list []int) (int, float64) {
    permutations := 0

    t0 := time.Now()
    for ! check_sorted(list) {
        list = Shuffle(list)
        permutations += 1
    }
    t1 := time.Now()
    duration := t1.Sub(t0)
    return permutations, duration.Seconds()
}


//Build_list returns a randomly shuffled list of numbers from 1 to len.
func Build_list(len int)[]int {
    list := make([]int, len)
    for i := 0; i < len; i++ {
        list[i] = i + 1
    }
    return Shuffle(list)
}


