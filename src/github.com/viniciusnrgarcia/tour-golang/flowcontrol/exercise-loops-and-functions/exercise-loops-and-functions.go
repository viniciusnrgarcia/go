package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	// eremos encontrar o número z para o qual z² é quase o x.

	z := float64(1)
	for i := 0; i <= 10; i++ {
		z = float64(i)
		z = (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	/*
	   Como forma de brincar com funções e loops, vamos implementar uma função de raiz quadrada: dado um número x, queremos encontrar o número z para o qual z² é quase o x.

	   x   = 16
	   z² ~=  4
	*/
	fmt.Println(Sqrt(16))

}
