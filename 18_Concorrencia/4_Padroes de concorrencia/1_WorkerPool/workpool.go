package main

import (
	"fmt"
	"time"
)

func main() {
	inicio := time.Now()
	//====
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)
	//Quanto mais goroutines adcionar mais rapido a execução será, obdecendo a capacidade de processamento da maquina
	//Funçao worker fica esperando tarefas chegarem para executar
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados) //com 3 go routines o tempo de execução reduziu em 50%

	//Conforme ocorre as iterações no laço, as tarefas são enviadas para a função worker que está aguardando o recebimento delas
	for i := 0; i < 45; i++ {
		tarefas <- i //enviando dados
	}
	//Após o canal de tarefas enviar os dados, deve ser fechado
	close(tarefas)

	//canal de resultados recebendo os dados enviados pelo canal de tarefas para a função worker
	for j := 0; j < 45; j++ {
		resultado := <-resultados //recebendo dados e armazenando na variavel
		fmt.Println(resultado)

	}
	fim := time.Now()

	duracao := fim.Sub(inicio)
	fmt.Println("Tempo de execução: ", duracao)

}

// A função worker é responsavel por executar as tarefas e envia-las para o canal de resultados
func worker(tarefas <-chan int, resultados chan<- int) { //tarefas recebe dados e resultados envia
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(numero int) int {
	if numero <= 1 {
		return numero
	}
	return fibonacci(numero-2) + fibonacci(numero-1)

}
