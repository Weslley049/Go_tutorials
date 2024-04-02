package main

import (
	"fmt"
)

type QuadradoType struct {
	base   float64
	altura float64
}

func (quadrado QuadradoType) calcularAreaQuadrado() float64 {
	return quadrado.base * quadrado.altura
}

type TrianguloType struct {
	base   float64
	altura float64
}

func (triangulo TrianguloType) calcularAreaTriangulo() float64 {
	return (triangulo.base * triangulo.altura) / 2
}

const PI float64 = 3.14159

// Definindo uma struct que utiliza a constante
type CirculoType struct {
	Raio float64
}

// Método para calcular a área do círculo
func (c CirculoType) CalcularAreaDoCirculo() float64 {
	return PI * (c.Raio * c.Raio)
}

func main() {

	quadradoInstancia := QuadradoType{60.7, 32.4}
	triangulo := TrianguloType{30.5, 20.6}

	circulo := CirculoType{
		Raio: 5.0,
	}

	var areaQuadrado = quadradoInstancia.calcularAreaQuadrado()
	var areaTriangulo = triangulo.calcularAreaTriangulo()

	fmt.Printf("A Base do quadrado é %vm²\n", areaQuadrado)

	fmt.Printf("A Base do triangulo é %vm²\n", areaTriangulo)

	// Calculando e exibindo a área do círculo
	fmt.Printf("Área do círculo: %.2f\n", circulo.CalcularAreaDoCirculo())

}
