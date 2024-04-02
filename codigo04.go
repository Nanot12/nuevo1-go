package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ContarPalabras(frase string) int {
	palabras := strings.Fields(frase)

	return len(palabras)
}

func main() {
	fmt.Print("Ingrese una frase: ")
	br := bufio.NewReader(os.Stdin)
	frase, err := br.ReadString('\n')

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //finalizar el programa con error
	}

	numPalabras := ContarPalabras(frase)

	fmt.Println("Número de palabras: ", numPalabras)

}
