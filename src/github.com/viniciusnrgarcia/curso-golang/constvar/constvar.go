package main

import (
	"fmt"
	math "math"
	"time"
)

const (
	a = 1
	b = 2
)

var (
	c = 3
	d = 4
)

func main() {
	time.Sleep(100)

	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduziada
	area := 3.2

	area = PI * math.Pow(raio, 2)
	fmt.Println("Area é", area)

	fmt.Println("Variaveis e constantes", a, b, c, d)

	// boolean
	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)

}
