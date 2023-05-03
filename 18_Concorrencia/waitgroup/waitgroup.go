package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitgroup sync.WaitGroup //var do tipo sync.WaitGroup

	waitgroup.Add(2) //Cria uma fila com um tamanho definido (2) de goroutines a serem executadas

	go func() {
		escrever("Olá Gophers!")
		waitgroup.Done() // quando finalizar execução decrementa -1 na fila de goroutines
	}()

	go func() {
		escrever("Golang é demais")
		waitgroup.Done() // quando finalizar execução decrementa -1 na fila de goroutines
	}()

	waitgroup.Wait() // bloqueia a execução até que todas as goroutines sejam executadas, no caso a fila seja = 0

	fmt.Println("Aquiiiiii")
}

func escrever(mensagem string) {
	for i := 1; i < 5; i++ {
		fmt.Println(mensagem)
		time.Sleep(time.Second)
	}
}
