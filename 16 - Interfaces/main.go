package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	// return (c.raio * c.raio) * math.Pi
	return math.Pow(c.raio, 2) * math.Pi
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func getArea(f forma) {
	fmt.Printf("A área da forma é %0.2f \n", f.area())
}

func main() {
	r := retangulo{10, 15}
	c := circulo{10}

	getArea(r)
	getArea(c)
}
