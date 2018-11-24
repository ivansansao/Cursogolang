package main

import (
	"fmt"
)

func main() {

	i := 1

	var p *int = nil
	p = &i // p recebe o endereço de i
	*p++   // Incrementa o valor contido na posição de memória
	i++

	// Go não possui aritmética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)

}
