package enderecos_test

import (
	"testing"
	"tipos-enderecos/enderecos"
)

// Struct de cenario de testes com dois estados
type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoEndereco(t *testing.T) {
	//Testando multiplos cenarios usando struct
	cenarioTeste := []cenarioTeste{
		{"Rua Dos Bobos", "Rua"},
		{"Avenida Principal", "Avenida"},
		{"Praça Central", "Tipo inválido"},
		{"Rodovia dos Bandeirantes", "Rodovia"},
		{"Travessa da Felicidade", "Travessa"},
		{"", "Tipo inválido"},
	}
	//Pecorrendo o struct de cenarioTeste e pegando cada elemento e enviando para a função TipoEndereco()
	for _, tipoEsperado := range cenarioTeste {
		tipoRecebido := enderecos.TipoEndereco(tipoEsperado.enderecoInserido)
		if tipoRecebido != tipoEsperado.retornoEsperado {
			t.Errorf("Esperava: %s e Recebeu: %s", tipoEsperado, tipoRecebido)
		}

	}

}
