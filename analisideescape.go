package main

import (
	"fmt"
	"time"
)

func countdown(seconds *int) {
	for *seconds > 0 {
		time.Sleep(time.Second * 1)
		*seconds -= 1
	}
}

func main() {
	count := 5
	//llamando al gorotine de valor inicial

	go countdown(&count) //referencia en memoria
	for count > 0 {
		time.Sleep(time.Millisecond * 500)
		//leer el valor de variable
		fmt.Println(count)
	}

}
