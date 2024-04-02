package main

import (
	"fmt"
	"time"
)

func mensaje(i int) {
	fmt.Println("sms ", i)
}

func main() {
	for i := 1; i <= 10; i++ {
		go mensaje(i)
	}

	time.Sleep(100 * time.Nanosecond)

	fmt.Println("Fin")
}
