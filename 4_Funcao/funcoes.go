package main

import "fmt"

//Função simples com parametro
func soma(n1, n2 int) int {
	return n1 + n2

}

//Função com retorno duplo
func resultadoCalculo(n1, n2 int) (int, int) {
	var soma = n1 + n2
	var sub = n1 - n2
	return soma, sub

}

func main() {
	fmt.Println(soma(10, 20))

	soma, sub := resultadoCalculo(30, 50)

	fmt.Println(soma, sub)

	//Somente a soma
	soma, _ = resultadoCalculo(30, 50)
	fmt.Println(soma)

	//Somente a subtração
	_, sub = resultadoCalculo(30, 50)
	fmt.Println(sub)

	//Função em variavel
	var f = func(nome string) string {

		return "olá " + nome + ", Seja bem vindo(a)"
	}
	funcaoEmVar := f("Caique")
	fmt.Println(funcaoEmVar)

}
