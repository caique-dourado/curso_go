package main

import "fmt"

// Interface
type animal interface {
	som() string
}

// struct
type cachorro struct{}
type gato struct{}

// Metodo com a mesma assinatura do metodo da interface
// Com isso, cachorro se torna automaticamente um animal
func (d cachorro) som() string {
	return "Au au au!"
}

func (g gato) som() string {
	return "Miauuuuu!"
}

func fazerSom(a animal) {
	fmt.Println(a.som())
}

func main() {
	//cao é uma instancia de cachorro
	cao := cachorro{}
	fazerSom(cao)

	gato := gato{}
	fazerSom(gato)

}
