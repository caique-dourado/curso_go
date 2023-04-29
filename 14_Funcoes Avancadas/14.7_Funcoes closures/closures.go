package main

import "fmt"

func main() {
	mensagem := "Ola mundo"

	saudacao := func() {
		fmt.Println(mensagem)

	}
	saudacao()

}
