package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
    // Struct em JSON
	c := cachorro{Nome: "Rex", Raca: "Dálmata", Idade: 3}

	cachorroEmJSONByte, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

    cachorroEmJSON := string(cachorroEmJSONByte)
	fmt.Println(cachorroEmJSON)
    fmt.Printf("%T\n", cachorroEmJSON)


    // Map em JSON
	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
    fmt.Printf("%T\n", bytes.NewBuffer(cachorro2EmJSON))
}