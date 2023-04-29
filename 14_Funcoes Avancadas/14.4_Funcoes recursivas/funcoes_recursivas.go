package main

import "fmt"

func fibonacci(numero uint) uint {
	if numero <= 1 {
		return numero
	}
	return fibonacci(numero-2) + fibonacci(numero-1)

}

func main() {
	n1 := uint(10)

	calculaFibonacci := fibonacci(n1)

	fmt.Println(calculaFibonacci)
}
