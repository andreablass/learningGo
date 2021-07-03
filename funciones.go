package main

import (
	"fmt"
	"math"
)

func areaRectangulo(a, b float64) float64 {
	return a * b
}

func areaTrapecio(a, b, c float64) float64 {
	return ((a + b) * c) / 2
}

func areaCirculo(a float64) float64 {
	return (math.Pi * a * a)
}

func main() {

	//Area Rectangulo b * h
	AreaRectangulo := areaRectangulo(5, 2)
	fmt.Println("Area rectangulo:", AreaRectangulo)

	//Area Trapecio ((Base mayor + base menor) * altura ) / 2
	AreaTrapecio := areaTrapecio(14, 20, 5)
	fmt.Println("Area Trapecio:", AreaTrapecio)

	//Area Circulo (PI * radio al cuadrado)
	AreaCirculo := areaCirculo(5)
	fmt.Println("Area Circulo:", AreaCirculo)
}
