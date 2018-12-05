package main

import "fmt"

func main() {

	// Cria uma slice de 10 cujo aponta para um array interno de 10.
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	// Atribuo outro slice de 10 cujo aponta para um array interno de 20.
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s)) // Imprime o slice, 10, 20.

	// Adiciona mais elementos no slice, preenchendo todas as vinte posições.
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// Não dá erro se adicionar mais slice que o array interno, ele cresce.
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

}
