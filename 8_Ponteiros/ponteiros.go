package main

import "fmt"

func main() {
	var n1 int = 200
	var ponteiro *int
	//Sempre usar o & para armazenar uma referência
	ponteiro = &n1
	//Exibe apenas a referencia de memoria
	fmt.Println(ponteiro)
	//Exibe o valor o valor (desreferenciação)
	fmt.Println(*ponteiro)
}
