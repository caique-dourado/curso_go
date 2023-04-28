package main

import "fmt"

func calc(n1 int, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2

	return
}

func main() {
	n1, n2 := 50, 30
	soma, subtracao := calc(n1, n2)
	fmt.Println("Valores passados:", n1, "e", n2)
	fmt.Println("Resultado soma:", soma, "\nResultado subtração", subtracao)

}
