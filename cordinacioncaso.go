package main

import (
	"fmt"
	"time"
)

func stingy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money += 10
	}
	fmt.Println("finaliza stingy")
}

func spendy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money -= 10
	}
	fmt.Println("finaliza spendy")
}

func main() {
	money := 100 //saldo inicial

	go stingy(&money)
	go spendy(&money)
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("saldo final: ", money)
}
