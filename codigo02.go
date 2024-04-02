// factorial
package main

import "fmt"

//funciones propias
func factorial(n int) int {
	//limite
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("El factorial de 5 es: ", factorial(5))
}
