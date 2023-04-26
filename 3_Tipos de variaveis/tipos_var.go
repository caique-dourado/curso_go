package main

import (
	"errors"
	"fmt"
)

func main() {
	//INTEIROS - São definido pelo numero de bytes
	var numero1 int8 = 8
	var numero2 int16 = 16
	var numero3 int32 = 32
	var numero4 int64 = 64
	var numero5 int = 200 // Pega o numeros de vytes de acordo com a arquitetura do PC
	fmt.Println(numero1, numero2, numero3, numero4, numero5)
	//UNSIGNED INT - INTEIROS SEM SINAIS
	var numero6 uint = 8
	fmt.Println(numero6)

	//FLOAT
	var numero7 float32 = 12.32
	var numero8 float64 = 12.327245646
	fmt.Println(numero7, numero8)

	//BOOLEAN
	var boleano1 bool = true
	var boleano2 bool //Valor zero, logo por default a saída é false
	var boleano3 bool = false
	fmt.Println(boleano1, boleano2, boleano3)

	//ERROR
	var err error
	fmt.Println(err)
	var errPerson error = errors.New("Error personalizado")
	fmt.Println(errPerson)
}
