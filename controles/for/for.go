package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { // Isso é o While de outras linguagens, se tirar o i++ fica infinito.
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 { // For convencional
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando... ")
	for i := 1; i <= 10; i++ {

		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Ímpar ")
		}

	}

	for {
		// Laço infinito
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}

}
