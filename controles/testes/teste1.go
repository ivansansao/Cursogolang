package main

import (
	"fmt"
)

func main() {

	var caixa string = "Laranja"

	for len(caixa) > 0 {

		fmt.Println(caixa[0:1])
		caixa = caixa[1:]

	}
}
