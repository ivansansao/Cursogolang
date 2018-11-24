package main

import (
	"fmt"
)

func main() {

	x := 2
	y := 3

	fmt.Println("Resultado =", x+y)

	type Cachorro struct {
		nome  string
		idade int
	}

	c1 := Cachorro{"Nina", 3}
	c2 := Cachorro{"Bio", 5}

	fmt.Println(c1.nome, c2.nome)

}
