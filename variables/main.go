package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Tipos de Declarações de variaveis VAR e CONST

	//var variavelModificada int
	const variavelNaoModificada string = "Aqui não posso mudar"

	//Tipos de Inteiros
	var intNum int = 10
	var intNum16 int16 = 10000
	var intNum32 int32 = 3000000
	var intNum64 int64 = 3000000000000000000

	fmt.Println(intNum)
	fmt.Println(intNum16)
	fmt.Println(intNum32)
	fmt.Println(intNum64)

	// Tipos de Float

	var floatNum32 float32 = 10.2020202020202020202022222222222222222222222222
	var floatNum64 float64 = 10.22222222222222222222222222222222222222222222222222222222222222222222222222222222222222222

	fmt.Println(floatNum32)
	fmt.Println(floatNum64)
	fmt.Println(intNum32)
	fmt.Println(intNum64)

	// String

	var MyString string = "Minha string"

	//A função len me mostra a quantidade de bits de um texto
	fmt.Println(len(MyString))

	//Para mostrar a quantidade de length de um array de char temos outra função
	fmt.Println(utf8.RuneCountInString(MyString))

	//Boolean
	var MyBoolean bool = false

	fmt.Println(MyBoolean)

	//Rune - Rune é um tipo de caractere que se utiliza com single quotes

	var MyRune rune = 'A'

	fmt.Println(MyRune)
}
