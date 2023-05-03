package main

import "fmt"

func main() {
	channel := make(chan string, 2) // cria um canal com espaço para 2 valores armazenados em buffer

	channel <- "Olá mundo!"     // envia o primeiro valor para o canal
	channel <- "Olá gophers..." // envia o segundo valor para o canal
	//Quando ocorre o primeiro envio, o canal não é bloqueado pois ainda
	//tem um valor em buffer a ser enviado

	//Recebe e mostra na tela os valores enviados pelo canal
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
