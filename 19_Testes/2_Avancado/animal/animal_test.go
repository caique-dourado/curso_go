package animal_test

import (
	"teste-avancado/animal"
	"testing"
)

func TestSomAnimal(t *testing.T) {
	t.Run("Cachoro", func(t *testing.T) {
		c := animal.Cachorro{}
		somRecebido := animal.FazerSom(c)
		somEsperado := "Au au au"
		if somRecebido != somEsperado {
			t.Fatalf("O som recebido: %s é diferente do som esperado: %s", somRecebido, somEsperado)

		}
	})

	t.Run("Gato", func(t *testing.T) {
		g := animal.Gato{}
		somRecebido := animal.FazerSom(g)
		somEsperado := "Miau"
		if somRecebido != somEsperado {
			t.Fatalf("O som recebido: %s é diferente do som esperado: %s", somRecebido, somEsperado)
		}
	})
}
