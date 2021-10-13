package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)
	// uint8 (1 byte - 0 à 256), uint16 (short 2 bytes), uint32 (int 4 bytes), uint64 (long 8 bytes)

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal
	// int8, int16, int32, int64
	i1 := math.MaxInt8
	fmt.Println("O valor máximo do int é", i1)

	i2 := math.MaxInt16
	fmt.Println("O valor máximo do int é", i2)

	i3 := math.MaxInt32
	fmt.Println("O valor máximo do int é", i3)

	i4 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i4)

	var i5 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i5))
	fmt.Println(i5)

	// Números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	var x2 float64 = 49.99
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(x2))

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	if bo {
		fmt.Printf("\nA variável %v é verdeira", bo)
	} else {
		fmt.Printf("\nA variável %v é falsa", bo)
	}

	// String
	s1 := "Ola meu nome é Vinicius"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// String com multiplas linhas
	s2 := `Olá
	meu
	nome é Vinici
	us`
	fmt.Println("O tamanho da string é", len(s2))
	fmt.Println(s2)

	// char - não possui o tipo
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)

	fmt.Println(`
	Variáveis
	A instrução var declara uma lista de variáveis, como em listas de argumentos de função, o tipo é o último passado.

	A declaração var pode estar num pacote ou a nível de função. Nós vemos ambos neste exemplo.`)
	var i int
	fmt.Println(i, c, python, java)

}

var c, python, java bool = true, true, true
