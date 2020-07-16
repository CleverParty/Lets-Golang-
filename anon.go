package main

import (
	"fmt"
)


func main(){
   
	anonTest := func() {
		fmt.Println("the anon function has been executed")
	}
	anonTest()
}