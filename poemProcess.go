package main

import (
	"fmt"
	"regexp"

	"gopkg.in/jdkato/prose.v2"
)

func main() {

	val, err := regexp.MatchString("fo.", "whatever")
	fmt.Println(val, err)
	rng := []int{2, 3, 4}
	for index, value := range rng {
		fmt.Printf("a[%d] = %d\n", index, value)
	}
	valueFunc := func() {
		fmt.Println("Anonymous func generated :")
	}
	valueFunc()
	doc, _ := prose.NewDocument("go is an interesting language:.")
	for _, token := range doc.Tokens() {
		fmt.Println(token.Text)
		//Tokenization - works good
	}
}
