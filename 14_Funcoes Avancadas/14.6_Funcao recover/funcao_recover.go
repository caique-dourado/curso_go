package main

import "fmt"

//recupera programa

func recupera() {
	if r := recover(); r != nil {
		fmt.Println("Programa recuperado")
	}
}

func calcMedia(n1, n2 float64) bool {

	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	// Sempre deve-se usar o defer ao recuperar uma execução
	defer recupera()
	panic("A média não pode ser 6")
}
func main() {
	aluno := calcMedia(6, 6)

	fmt.Println(aluno)
}
