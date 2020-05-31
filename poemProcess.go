package main

import (
	"fmt"
	"math/rand"

	"gopkg.in/jdkato/prose.v2"
)

func main() {

	a := [3]int{1, 2, 3} // first method
	var b [3]string      // second method of declaring the arrays
	fmt.Println(a)
	a[0] = rand.Intn(100)
	a[1] = rand.Intn(100)
	a[2] = rand.Intn(100)
	b[0] = "India"
	b[1] = "Canada"
	b[2] = "Japan"
	for index, value := range a {
		fmt.Printf("a[%d] = %d\n", index, value)
	}
	doc, _ := prose.NewDocument("go is an interesting language:.")
	for _, token := range doc.Tokens() {
		fmt.Println(token.Text)
		//Tokenization - works good
	}
}
