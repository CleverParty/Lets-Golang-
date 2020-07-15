package main

import (
	"fmt"

	"gopkg.in/jdkato/prose.v2"
)

func main() {

	rng := 3
	for index, value := range rng {
		fmt.Printf("a[%d] = %d\n", index, value)
	}
	doc, _ := prose.NewDocument("go is an interesting language:.")
	for _, token := range doc.Tokens() {
		fmt.Println(token.Text)
		//Tokenization - works good
	}
}
