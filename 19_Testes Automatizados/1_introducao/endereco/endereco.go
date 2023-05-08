package endereco

import "strings"

func TipoEndereco(endereco string) string {
	tipos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoValido := false
	enderecoMinusculo := strings.ToLower(endereco)
	enderecoRecebido := strings.Split(enderecoMinusculo, " ")
	primeiraPalavra := enderecoRecebido[0]

	for _, tipo := range tipos {
		if tipo == primeiraPalavra {
			enderecoValido = true
		}
	}

	if enderecoValido {
		return strings.Title(primeiraPalavra)
	}
	return "Endereço inválido"

}
