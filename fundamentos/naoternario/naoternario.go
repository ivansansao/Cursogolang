package main

import (
	"fmt"
)

// Não tem operador ternário em Go.
func obterResultado(nota float64) string {

	if nota >= 6 {
		return "Aprovado"
	}

	return "Reprovado"

}

func main() {

	fmt.Println("Resutado:", obterResultado(6.2))

}
