package main

import "fmt"

func main() {
	// Criando array
	var array1 [4]int
	// Populando array
	array1[0] = 1
	array1[1] = 200
	array1[2] = 75
	array1[3] = 27

	fmt.Println(array1)

	//Populando inline / Partial
	array1 = [4]int{45, 46, 64, 8}
	fmt.Println(array1)

	//Criando um Slice
	var slice []string
	slice = []string{"Davi", "Elias", "Jó", "Tiago", "Daniel"}
	fmt.Println(slice)
	slice = append(slice, "Moisés")

	fmt.Println(slice)

	//Por debaixo dos panos
	var slice2 = make([]float32, 5, 8)

	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	//Extourando a capacidade, o Golang dobra ela.
	slice2 = append(slice2, 4.6)
	slice2 = append(slice2, 4.7)
	slice2 = append(slice2, 8.8)
	slice2 = append(slice2, 9.9)

	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	fmt.Println(slice2)

}
