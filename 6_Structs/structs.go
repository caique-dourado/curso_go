package main

import "fmt"

// Funciona basicamente como uma classe, ela é um modelo de uma estrutura
type usuario struct {
	nome  string
	idade uint16
	sexo  string
}

type endereco struct {
	logradouro string
	numero     uint8
}

type usuario2 struct {
	nome     string
	idade    uint16
	sexo     string
	endereco endereco
}

func main() {
	//Cada variavel no struct é referenciada por sua posição
	//Mas da forma abaixo (nomeada), a posição não importa
	var user = usuario{idade: 26, nome: "Caique", sexo: "Masculino"}
	fmt.Println(user)

	fmt.Println("\nRecuperando um valor por vez")
	fmt.Println(user.nome)
	fmt.Println(user.idade)
	fmt.Println(user.sexo)

	fmt.Println("\nStruct dentro de Struct")
	var enderecoCompleto = endereco{logradouro: "Jerusalem - Israel", numero: 100}
	var user2 = usuario2{nome: "Davi", idade: 70, sexo: "Masculino", endereco: enderecoCompleto}
	fmt.Println(user2)
}
