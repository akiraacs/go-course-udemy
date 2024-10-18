package main

import "fmt"


type animal interface {
	falar() string
}

type cachorro struct {
	nome string
	raca string
}

func (c cachorro) falar() string {
	return "Au Au"
}

type gato struct{
    nome string
    raca string
}

func (g gato) falar() string {
    return "Miau"
}

func interagirComAnimal(animal animal, nomeAnimal string) {
    fmt.Printf("%s está falando: %s\n", nomeAnimal, animal.falar())
}

func main() {
    animal1 := gato{nome: "Xanim", raca: "Siames"}
    animal2 := cachorro{nome: "Tobi", raca: "Pitbull"}

    fmt.Println(animal1.nome, animal1.raca, animal1.falar())
    fmt.Println(animal2.nome, animal2.raca, animal2.falar())

    interagirComAnimal(animal1, animal1.nome)
}
