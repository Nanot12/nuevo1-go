//i need a function main can 
//the algorithm 2.10 incrementing and decrementing
//with the integer n <- 0
// p  | q 
// --- | ---
// interger temp | interger temp 
// p1: do K times| q1 : do K times
// p2: temp <- n | q2 : temp <-n 
// p3: n <- temp | q3 : n <- temp -1

package main

import (
    "fmt"
    "sync"
)

var (
    n    int
    lock sync.Mutex
)

func P() {
    for i := 0; i < 5; i++ { // Assuming K = 5
        lock.Lock()
        n++
        fmt.Printf("incrementa n: %d\n", n)
       lock.Unlock()
    }
}

func Q() {
    for i := 0; i < 5; i++ { // Assuming K = 5
        lock.Lock()
        n--
        fmt.Printf("decrementa n: %d\n", n)
        lock.Unlock()
    }
}

func main() {
    go P()
    go Q()

    // Wait for goroutines to finish
    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        defer wg.Done()
        P()
    }()
    go func() {
        defer wg.Done()
        Q()
    }()
    wg.Wait()
}
