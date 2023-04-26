package main

import "fmt"

func main() {
	//Forma 1
	var varivael1 string = "Variavel 1"
	fmt.Println(varivael1)

	//Forma 2 - inferencia de tipo
	variavel2 := "Variavel 2"
	fmt.Println(variavel2)

	//Forma 3 - declarando multiplas variaveis

	var (
		nome      string = "Caique"
		sobrenome string = "Dourado"
		sexo      string = "Masculino"
	)

	//Multiplos valores
	nome, sobrenome, sexo = "Ana", "Cristina", "Feminino"

	fmt.Println(nome, sobrenome, sexo)

	const nacionalidade = "Brasil"
	fmt.Println(nome, sobrenome, sexo, nacionalidade)

}
