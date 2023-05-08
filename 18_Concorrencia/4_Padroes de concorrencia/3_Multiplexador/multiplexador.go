package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexador(escrever("Olá, Gopher"), escrever("Golang é muito legal!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexador(cnEntrada1, cnEntrada2 <-chan string) <-chan string {
	cnSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-cnEntrada1:
				cnSaida <- mensagem

			case mensagem := <-cnEntrada2:
				cnSaida <- mensagem
			}
		}
	}()

	return cnSaida

}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Recebido %s", texto)

			time.Sleep(time.Second * 1)
		}
	}()

	return canal
}
