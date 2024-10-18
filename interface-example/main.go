package main

import (
	"fmt"
	"math"
)

// Interface que define o método de cálculo da área
type Forma interface {
	area() float64
}

// Tipo "Retangulo"
type Retangulo struct {
	largura, altura float64
}

func (r Retangulo) area() float64 {
	return r.largura * r.altura
}

// Tipo "Circulo"
type Circulo struct {
	raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

// Função que recebe uma "Forma" e imprime sua área
func calcularArea(f Forma) {
	fmt.Printf("A área é: %.2f\n", f.area())
}

func main() {
	r := Retangulo{largura: 5, altura: 3}
	c := Circulo{raio: 2.5}

	calcularArea(r)
	calcularArea(c)
}
