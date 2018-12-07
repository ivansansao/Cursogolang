package main

import "fmt"

func main() {

	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)

	fmt.Println(s1, s2)

	// Para provar que um slice aponta para um array interno.
	s1[0] = 7

	fmt.Println(s1, s2)

}
