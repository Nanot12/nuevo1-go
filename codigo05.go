// Algoritmo trivial concurrente
package main

import (
	"fmt"
	"time"
)

//variable global

var n int

func P() {
	k1 := 1
	n = k1
}

func Q() {
	k2 := 2
	n = k2
}

func main() {
	n = 0

	go P()
	go Q()

	time.Sleep(time.Nanosecond * 10) //pausa
	fmt.Println("Valor final de n= ", n)
}
