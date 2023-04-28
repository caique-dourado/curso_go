package main

import "fmt"

func main() {
	//For convencional
	fmt.Println("FOR CONVENCIONAL")

	n1 := 0
	for n1 < 10 {
		fmt.Println("Valor atual de i:", n1)
		n1++
	}

	// For com init
	fmt.Println("\nFOR INIT")

	n2 := 0
	for j := n2; j < 10; j++ {
		fmt.Println("O valor atual do i:", j)
	}

	//Percorrendo arrays / slice / maps é do mesmo jeito
	//Basta usar a palavra reservada <range>
	fmt.Println("\nPERCORRENDO ARRAYS / SLICE / MAPS")

	arrayNomes := [3]string{"João", "Davi", "Tiago"}

	for indice, nome := range arrayNomes {
		fmt.Println(indice, nome)
	}

	//Caso não queira o indice, basta usar o (_) underscore
	fmt.Println("\nPERCORRENDO ARRAYS / SLICE / MAPS - SEM O INDICE")

	//Essa forma é muito comum
	for _, nome := range arrayNomes {
		fmt.Println(nome)
	}

}
