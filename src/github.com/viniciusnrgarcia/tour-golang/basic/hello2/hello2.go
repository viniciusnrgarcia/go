package main

import (
	"fmt"
	"rsc.io/quote"
	"time"
)

func main() {
	time.Sleep(100)
	fmt.Println("Hello, World golang!")
	fmt.Println("Waiting 1s")
	time.Sleep(1 * time.Second)
	fmt.Println("Show!")

	fmt.Println(quote.Go())

	fmt.Println(quote.Hello())

	fmt.Println(quote.Opt())
}
