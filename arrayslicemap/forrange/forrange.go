package main

import "fmt"

func main() {

	numeros := [...]int{36, 40, 15, 27, 89}

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}

}
