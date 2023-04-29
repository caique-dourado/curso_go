package main

import "fmt"

func main() {
	func() {
		fmt.Println("Executando uma função anonima")
	}()

	//Quando se tem retorno e deseja usa-lo, recomenda-se armazenar a funcao anonima em uma VAR
	funcao := func(nome string) string {
		return "Olá, Sr(a): " + nome
	}("Caique Dourado")

	fmt.Println(funcao)

}
