package main

import (
	"fmt"

	"gopkg.in/jdkato/prose.v2"
)

func main() {

	a := [3]int{1, 2, 3}
	fmt.Println(a)
	for index, value := range a {
		fmt.Printf("a[%d] = %d\n", index, value)
	}
	doc, _ := prose.NewDocument("go is an interesting language:.")
	for _, token := range doc.Tokens() {
		fmt.Println(token.Text)
		//Tokenization - works good
	}
}
