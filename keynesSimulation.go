package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func ratings(x int, y int) int {
	return x + y
}

func process(a int, b int) int {
	ele := a + b
	return ele
}

func main() {

	a := [3]int{1, 2, 3} // first method
	var b [3]string      // second method of declaring the arrays
	fmt.Println(a)
	var avg, sumVal, storeDif int
	var y, j int
	a[0] = rand.Intn(100)
	a[1] = rand.Intn(100)
	a[2] = rand.Intn(100)
	b[0] = "India"
	b[1] = "Canada"
	b[2] = "Japan"
	begin := time.Now()
	for index, value := range a {
		sumVal += value
		fmt.Printf("a[%d] = %d\n", index, value)
		if index == 2 {
			fmt.Println(sumVal)
			avg = sumVal / 3
			for index, innerValue := range a { // this loop can be added to a function -- > future enhancement
				fmt.Println((2 * avg) / 3)
				storeIndex := (2*avg)/3 - innerValue
				f := float64(storeIndex)
				dif := math.Abs(f) //value - innerValue
				y = int(dif)
				storeDif = y
				if y <= storeDif {
					storeDif = y
					fmt.Printf("lowest difference := %d\n", value)
					j = index
				}
				val := process(storeIndex, storeDif)
				fmt.Println("[val] value is :", val)
				storeDif = int(dif)
				fmt.Printf("the differences for each rating[%d] from 2/3*mean = %d \n", j, storeDif)
				fmt.Printf("\nthe winner is player no: %d \n", index)
			}
		}
	}
	total := time.Since(begin)
	fmt.Printf("The winner was calculated in : %v \n", total)
	fmt.Printf("The contestant with closest value picked is := %d \n", j)
	// for i := 0; i < 10; i++ {
	//  	a[i] = rand.Intn(100)
	//	    fmt.Println(a[i])
	// }python3 -m venv examplep
}
