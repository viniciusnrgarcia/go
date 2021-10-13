package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	a()

}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
