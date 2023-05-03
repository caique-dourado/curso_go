package main

import "fmt"

func main() {
	channel := make(chan string)
	nomes := [5]string{"Davi", "José", "João", "Eliú", "Jó"}

	go saudacao(nomes, channel)

	for nome := range channel {

		fmt.Printf("Olá %s, seja bem vindo(a)\n", nome)
	}

}

func saudacao(nomes [5]string, channel chan string) {

	for _, nome := range nomes {
		channel <- nome

	}

	close(channel)

}
