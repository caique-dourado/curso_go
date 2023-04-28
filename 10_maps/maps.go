package main

import "fmt"

func main() {
	//Map simples
	usuario := map[string]string{
		"nome":      "Caique",
		"sexo":      "Masculino",
		"profissao": "Eng. Soft",
	}

	fmt.Println(usuario)

	//Printando valores por chave
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["profissao"])

	//map dentro de map
	usuario2 := map[string]map[string]string{
		"caique": {
			"nome_completo": "Caique Dourado",
			"nacionalidade": "Brasileiro",
			"profissao":     "Engenheiro de Software",
		},
	}

	fmt.Println(usuario2)

}
