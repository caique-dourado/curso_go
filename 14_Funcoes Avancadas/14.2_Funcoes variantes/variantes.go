package main

import "fmt"

func somatorio(numeros ...int) int {
	total := 0
	for _, numero := range numeros {

		total += numero

	}

	return total
}

func main() {
	soma := somatorio(1, 4, 6, 7, 8, 65, 54, 54, 12, 4, 5)

	fmt.Println(soma)
}
