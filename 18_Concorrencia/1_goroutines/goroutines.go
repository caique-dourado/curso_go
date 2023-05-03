package main

import (
	"fmt"
	"time"
)

func main() {
	//go-> inicia a concorrencia, fazendo com que aos poucos cada parte das tarefas vão sendo executadas
	go escrever("Olá Gophers!")
	escrever("Esse é o Golang")

}

func escrever(mensagem string) {
	for {
		fmt.Println(mensagem)
		time.Sleep(time.Second)
	}
}
