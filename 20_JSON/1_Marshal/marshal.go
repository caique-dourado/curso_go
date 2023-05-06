package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// criando um struct cachorro
type cachorro struct {
	Nome  string `json:"nome"` //definindo o padrão das minhas chaves
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	//Criando uma instancia de cachorro
	c := cachorro{"Rex", "Dalmata", 10}
	//Retornando um slice uint8  (bytes)
	cachorroJson, erro := json.Marshal(c)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroJson)
	//convertendo o slice em um json entendivel

	cachorroBuffer := bytes.NewBuffer(cachorroJson)

	fmt.Println(cachorroBuffer)

}
