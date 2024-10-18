package main

import (
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
	// JSON para struct
	cachorroEmJSON := `{"nome":"Rex", "raca":"Dálmata", "idade":3}`

	var c cachorro
	err := json.Unmarshal([]byte(cachorroEmJSON), &c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)
	fmt.Printf("%T\n", c)

	// JSON para map
	cachorro2EmJSON := `{"nome":"Toby", "raca":"Poodle"}`

	c2 := make(map[string]string)
	err = json.Unmarshal([]byte(cachorro2EmJSON), &c2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c2)
	fmt.Printf("%T\n", c2)
}
