package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade string `json:"idade"` //alterei o tipo para poder usar o map tambem
}

func main() {
	cachorroJson := `{"nome":"Princesa", "raca":"Poodle", "idade":"10"}`

	c := cachorro{}

	if erro := json.Unmarshal([]byte(cachorroJson), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	//usando um map
	cachorroJson2 := `{"nome":"Rex", "raca":"Husky", "idade":"2"}`
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorroJson2), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)
}
