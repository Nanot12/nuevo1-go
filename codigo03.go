// Ordenamiento de números
package main

import (
	"fmt"
	"sort"
)

func OrdenarNumeros(numeros []int) []int {
	sort.Ints(numeros)
	return numeros
}

func main() {
	//declaración de slice de enteros
	listaNumeros := []int{5, 3, 8, 1, 2}

	fmt.Println("Números oredenados: ", OrdenarNumeros(listaNumeros))
}
