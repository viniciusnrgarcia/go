package main

import "fmt"

/*
For é o "while" de Go
Nesse ponto, você pode tirar as vírgulas: o while do C é escrito com for em Go.
*/
func main() {
	sum := 1
	for sum < 100 {
		fmt.Println(sum)
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)
}
