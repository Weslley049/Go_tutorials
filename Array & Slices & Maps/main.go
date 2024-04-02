package main

import "fmt"

func main() {

	//Instância de Array
	var intArr [3]int32

	intArr[1] = 123

	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	//Instância de Array com valor direto

	intArray := [3]int32{1, 2, 3}

	fmt.Println(intArray)

	//Instância de Array sem tamanho definido

	intArraySemTamanhoDefinido := [...]int32{1, 2, 3, 4, 5}

	fmt.Println(intArraySemTamanhoDefinido)

	//Adição de novos itens no array

	intArrNovosItens := [...]int32{1, 2, 3}

	fmt.Println(intArrNovosItens)

	var intSlice []int32 = []int32{4, 5, 6}

	//Capacidade e Length de um array

	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7)

	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	//Adicionando outro array ao array original

	var intSlice2 []int32 = []int32{8, 9}

	intSlice = append(intSlice, intSlice2...)

	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)

	fmt.Println(intSlice3[2])

	//Maps

	var myMap map[string]uint8 = make(map[string]uint8)

	fmt.Println(myMap)

}
