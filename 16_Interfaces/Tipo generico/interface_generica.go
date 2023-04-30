package main

import "fmt"

func main() {
	//interface{} da flexibilidade ao montar um map que aceita valores dinamicos
	usuario := map[string]interface{}{
		"nome":  "Super Mario",
		"idade": 50,
		"sexo":  "M",
		"logradouro": map[string]interface{}{
			"rua":    "Rua da felicidade",
			"numero": 200,
			"cidade": "Mario World",
		},
	}

	fmt.Println(usuario)
}
