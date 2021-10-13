package main

import "fmt"

func main() {
	i, j := 42, 2701

	fmt.Println(j)
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i
	fmt.Println(*p)
	fmt.Println(p)

	p = &j // point to j
	fmt.Println(j)
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	x := 10
	y := &x

	fmt.Println(&x)
	//fmt.Println(&y)

	*y++
	fmt.Printf("%T, %T \n", x, y)
	fmt.Printf("%T, %v \n", x, x)
	fmt.Println(y)

	w := 11
	recebeValor(w)
	fmt.Println(w)

	z := 11
	recebePonteiro(&z)

	fmt.Println(z)
}

func recebeValor(x int) {
	x++
	fmt.Println("Na função:", x)
}

func recebePonteiro(x *int) {
	*x++
	fmt.Println("Na função:", *x)
}
