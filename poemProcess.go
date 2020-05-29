package main

import (
	"fmt"

	"gopkg.in/jdkato/prose.v2"
)

func main() {

	var a [10]int
	fmt.Println(a)
	doc, _ := prose.NewDocument("go is an interesting language:.")
	for _, token := range doc.Tokens() {
		fmt.Println(token.Text)
		//Tokenization - works good
	}
}
