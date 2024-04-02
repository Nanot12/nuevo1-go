package main

import "fmt"

func P() {
	for {
		fmt.Println("P line1 CS")
		fmt.Println("P line2 CS")
	}
}

func Q() {
	for {
		fmt.Println("Q line1 NCS")
		fmt.Println("Q line2 NCS")
		for turn != 2 {
			//esperando
		}
		turn = 1
	}
}

func main() {
	go P()
	Q()
}
