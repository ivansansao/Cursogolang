package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// Números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// Sem sinal (u é unsigned)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// Com sinal... int8 int16 int32 int64
	i1 := math.MaxFloat64
	fmt.Println("O valor máximo de int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // Representa uma mapeamento de uma tabela Unicode
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// Números reais (fload32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é,", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	// Boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// String
	s1 := "Olá meu nome é Leo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// String com múltiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Leo`

	fmt.Println("O tamanho da string é", len(s2))

	// Char não tem!

}
