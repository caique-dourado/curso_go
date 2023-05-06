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
	//convertendo um json em struct
	cachorroJson := `{"nome":"Princesa", "raca":"Poodle", "idade":"10"}`

	c := cachorro{}
	//isso []byte(cachorroJson) é o mesmo que bytes.NewBuffer(cachorroJson)
	if erro := json.Unmarshal([]byte(cachorroJson), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	//conertendo um json em map
	cachorroJson2 := `{"nome":"Rex", "raca":"Husky", "idade":"2"}`

	//criando o map
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorroJson2), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)
}
