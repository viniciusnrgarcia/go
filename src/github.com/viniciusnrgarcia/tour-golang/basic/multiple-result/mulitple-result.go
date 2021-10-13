package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func calculateAvailable(x, y int) int {
	return y + x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	result := calculateAvailable(1, 4)

	fmt.Println(result)
}
