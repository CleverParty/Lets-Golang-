package main

import (
	"fmt"
)

func anonFunc() {
	fmt.Println("the function which accepts other functions as arguements has been entered")
	// secondClass := func() {}

}

func secondClass() {
	fmt.Println("global status check")
}

func testerFunc(a func()) {

	fmt.Println("the Tester function has been entered")
}

func main() {

	firstClass := func(a, b int) int {
		value := a * b
		return value
	}
	anonTest := func() {
		fmt.Println("the anon function has been executed")
	}
	anonTest()
	fmt.Println("the value returned by first class func is: ", firstClass(12, 34))
	secondClass()
}
