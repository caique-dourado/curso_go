package animal

type animal interface {
	som() string
}

type Cachorro struct{}

type Gato struct{}

func (c Cachorro) som() string {
	return "Au au au"
}

func (g Gato) som() string {
	return "Miau"
}

func FazerSom(a animal) string {

	return a.som()

}
