package main

import "fmt"

func main() {
	fmt.Println("Operadores do go")
	fmt.Println("(+)SOMA")
	fmt.Println("(-)SUBTRAÇÃO")
	fmt.Println("(*)MULTIPLICAÇÃO")
	fmt.Println("(/)DIVISÃO")
	fmt.Println("(%)RESTO")

	//OPERAÇÕES SÓ SÃO REALIZADAS SE POSSUIREM O MESMO TIPO
	var n1 int16 = 20
	var n2 int16 = 25
	var soma = n1 + n2
	fmt.Println("A soma de n1 + n2 = ", soma)
	fmt.Println("\nOperadores unários")
	fmt.Println("(+=) -> N1 = N1 + (VALOR)")
	fmt.Println("(-=)-> N1 = N1 - (VALOR)")
	fmt.Println("(*=)-> N1 = N1 * (VALOR)")
	fmt.Println("(/=)-> N1 = N1 / (VALOR)")

}
