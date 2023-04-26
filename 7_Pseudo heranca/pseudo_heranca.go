package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
	sexo  string
}

type estudante struct {
	pessoa    //chamando o struct diretamente, pega todos os atributos
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{nome: "Elias", idade: 75, sexo: "Masculino"}
	estd1 := estudante{pessoa: p1, curso: "Teólogo", faculdade: "UNDB"}

	fmt.Println(estd1)
}
