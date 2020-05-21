package main

import (
	"fmt"

	"gopkg.in/jdkato/prose.v2"
)

func main() {
	doc, _ := prose.NewDocument("This is a sentence.")
	for _, token := range doc.Tokens() {
		fmt.Println(token.Text)
		//Tokenization - works good
	}
}
