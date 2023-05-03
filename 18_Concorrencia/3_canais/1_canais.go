package main

import "fmt"

func main() {
	channel := make(chan string) // cria um canal que só recebe dados do tipo string

	go escrever("olá mundo", channel) // Goroutine
	mensagem := <-channel             // clausula que recebe o dado do canal
	fmt.Println(mensagem)

}

func escrever(texto string, channel chan string) {
	channel <- texto // clausula que envia um dado para o canal
	close(channel)   // fecha o canal
}
