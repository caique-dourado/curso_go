package main

import "teste-avancado/animal"

func main() {
	c := animal.Cachorro{}
	g := animal.Gato{}

	animal.FazerSom(c)
	animal.FazerSom(g)

}
