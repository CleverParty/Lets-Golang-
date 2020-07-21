package main

import (
	"fmt"
	"math"
	"time"
)

func ratings(x int, y int) int {
	return x + y
}

func process(a int, b int) int {
	ele := a + b
	return ele
}

func run() string {
	return "This function execution will be treated as the canary in the coal mine."
}

func main() {

	var avg, sumVal, storeDif int
	var y, j int
	begin := time.Now()
	players := [3]int{1, 6, 3}
	fmt.Println(process(23, 78))
	for index, value := range players {
		sumVal += value
		fmt.Printf("a[%d] = %d\n", index, value)

		avg = sumVal / 3
		for index, innerValue := range players { // this loop can be added to a function -- > future enhancement
			storeIndex := (2*avg)/3 - innerValue
			f := float64(storeIndex)
			dif := math.Abs(f) //value - innerValue
			y = int(dif)
			storeDif = y
			if y <= storeDif {
				storeDif = y
				fmt.Printf("2/3*mean value is : %d\n", ((2 * avg) / 3))
				fmt.Printf("lowest difference := %d\n", value)
				j = index
			}
			val := process(storeIndex, storeDif)
			fmt.Println("[val] value is :", val)
			storeDif = int(dif)
			fmt.Printf("the differences for each rating[%d] from 2/3*mean = %d \n", j, storeDif)
			fmt.Printf("\nthe winner is player no: %d \n", index)
			fmt.Println(storeDif)
		}
	}
	total := time.Since(begin)
	fmt.Printf("The winner was calculated in : %v \n", total)
	fmt.Printf("The contestant with closest value picked is := %d \n", j)
	// for i := 0; i < 10; i++ {
	//  	a[i] = rand.Intn(100)
	//	    fmt.Println(a[i])
	// }python3 -m venv
}
