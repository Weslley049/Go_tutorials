package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func numeros() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}

	wg.Done()
}

func letras() {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 230)
	}

	wg.Done()
}

func main() {
	wg.Add(1)
	go numeros()

	wg.Add(1)
	go letras()

	wg.Wait()
	fmt.Println("Fim da execução")
}
