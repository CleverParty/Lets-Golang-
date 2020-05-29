package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var a [10]int
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a[i] = rand.Intn(100)
		fmt.Println(a[i])
	}
}
