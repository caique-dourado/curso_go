package main

import "fmt"

var n int

func init() {
	fmt.Println("Executou o init")
	n = 50
}
func main() {
	fmt.Println("Iniciou a main")
	fmt.Println(n)
	fmt.Println("Fim main")
}
