package main

import "fmt"

func inverteNumero(numero *int) {
	//O uso do ponteiro obtem o endereço da memoria altera o seu valo
	*numero = *numero * -1
}

func main() {
	n1 := 27
	fmt.Println("Valor de n1 antes de inverter:", n1)
	//&n1 -> está passando o endereço de memoria de n1, e não o valor 27
	inverteNumero(&n1)
	fmt.Println("Valor de n1 após chamar a funcao inverteNumero:", n1)

}
