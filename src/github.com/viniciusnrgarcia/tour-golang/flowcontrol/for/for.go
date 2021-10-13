package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	sum := 0
	t := 1000000000
	for i := 0; i < t; i++ {
		sum += i
	}
	fmt.Println(sum)

	elapsed := time.Since(start)
	log.Printf("\ntook %s to process %v", elapsed, t)

}

// func timeTrack(start time.Time, name string) {
// 	elapsed := time.Since(start)
// 	log.Printf("%s took %s", name, elapsed)
// }
