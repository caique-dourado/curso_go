package main

import "fmt"

func conexaoBanco() string {
	//O defer atrasa até o ultimo momento a execução da linha de codigo
	//No caso até um momento antes de executar o return
	defer fmt.Println("Fechando conexao com o banco")
	fmt.Println("Abrindo conexao com o banco")

	for i := 1; i < 10; i++ {
		fmt.Println("Fazendo requisição ao banco")
	}

	return string("Requisição conluida")
}

func main() {
	con := conexaoBanco()

	fmt.Println(con)
}
