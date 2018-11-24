package main

import (
	"fmt"
)

func compras(trab1, trab2 bool) (bool, bool, bool) {

	comprarTv50 := trab1 && trab2    // AND
	comprarTv32 := trab1 != trab2    // XOR (Negação)
	comprarSorvete := trab1 || trab2 // OR

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {

	tv50, tv32, sorvete := compras(true, false)

	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
