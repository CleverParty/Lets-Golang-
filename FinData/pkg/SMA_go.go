package main	

import (
	"fmt"
)

type vertex struct { // A struct is a collection of fields.
	X int
	Y int
}

func main() {

	i, j := 10, 11
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	point := &i

	fmt.Println(*point)
	*point = 13
	fmt.Println(i)

	point = &j
	fmt.Println(*point)
	*point = 14

	// struct testing :
	fmt.Println("Struct Testing")
	v := vertex{57481, 3}
	v.X = 635261
	fmt.Println(v.X)
	// pointers can also be made to structs

	point2 := &v
	point2.X = 404
	fmt.Println(v.X)

	fmt.Println(primes)
	// can declare and initialize the array at any poitn in the program
	var array [1]string
	array[0] = "canary test is working"

	fmt.Println(array)

}
