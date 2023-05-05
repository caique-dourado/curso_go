package enderecos

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TipoEndereco(endereco string) string {
	enderecoValido := false
	tiposEnderecos := []string{"rua", "avenida", "rodovia", "travessa"}
	enderecoMin := strings.ToLower(endereco)
	enderecoSplit := strings.Split(enderecoMin, " ")
	primeiraPalavra := enderecoSplit[0]
	for _, tipo := range tiposEnderecos {
		if tipo == primeiraPalavra {
			enderecoValido = true
		}
	}

	if enderecoValido {
		//converte a primeira letra em maiusculo
		caser := cases.Title(language.BrazilianPortuguese)
		return caser.String(primeiraPalavra)

	}

	return "Tipo inválido"
}
