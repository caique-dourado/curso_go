package main

import "fmt"

func main() {
	n1 := 15

	if n1 >= 25 {
		fmt.Println(n1, "É maior que 25")
	} else {
		fmt.Println(n1, "Não é maior que 25")
	}

	//estrtura do if init

	if n2 := n1; n2 > 10 {
		fmt.Println(n2, "É maior que 10")
	} else {
		fmt.Println(n2, "Não é maior que 10")
	}

}
