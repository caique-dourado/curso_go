package main

import "fmt"

// Criando um struct
type usuario struct {
	nome  string
	idade uint
}

// Criando um metodo relacionado ao struct [usuario]
func (u usuario) saudacao() {
	fmt.Printf("Olá Sr(a) %s, seja bem vindo(a)\n", u.nome)
}

// Criando um metodo fazer aniversario

func (u *usuario) fazerAnivesario() {
	u.idade++
	fmt.Printf("Parabéns sr(a) %s\nVocê completou %d anos", u.nome, u.idade)
}

//DICA PRINTF()
//%s -> pega string
//%d -> pega inteiros
//%g -> pega float
//%p -> pega ponteiros

func main() {
	usuario1 := usuario{"Caique", 26}
	usuario1.saudacao()
	usuario1.fazerAnivesario()
}
