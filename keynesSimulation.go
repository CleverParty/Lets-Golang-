package main

import (
	"fmt"
	"math/rand"
)

func main() {

	a := [3]int{1, 2, 3} // first method
	var b [3]string      // second method of declaring the arrays
	fmt.Println(a)
	var avg, sumVal int
	a[0] = rand.Intn(100)
	a[1] = rand.Intn(100)
	a[2] = rand.Intn(100)
	b[0] = "India"
	b[1] = "Canada"
	b[2] = "Japan"
	for index, value := range a {
		sumVal += value
		fmt.Printf("a[%d] = %d\n", index, value)
		if index == 2 {
			avg = sumVal / 3
		}
	}
	fmt.Println(avg)
	// for i := 0; i < 10; i++ {
	//  	a[i] = rand.Intn(100)
	//	    fmt.Println(a[i])
	// }
	fmt.Println("The ratings of the contestants %v:", a)
}
