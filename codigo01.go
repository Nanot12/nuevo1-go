// Suma de pares
package main

import "fmt"

var dato int //declaración de variable

//funciones propias
func SumaNumerosPares(n int) int {
	suma := 0 //declarando, definiendo el tipo según el valor asignado

	for i := 0; i <= n; i += 2 {
		suma += i
	}

	return suma
}

func main() {

	fmt.Println(SumaNumerosPares(10))
}
