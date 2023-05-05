package main

import (
	"fmt"
	"tipos-enderecos/enderecos"
)

func main() {
	endereco := enderecos.TipoEndereco("Rodovia dos imigrantes")
	fmt.Println(endereco)
}
