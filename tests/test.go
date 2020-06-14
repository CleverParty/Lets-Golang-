package main

import (
  "fmt"

  "github.com/skratchdot/open-golang/open"
)

type Node struct{
  Value int
  left *Node
  right *Node
}



func main() {
  one := &Node{
    Value : 10 ,
  }
  // var val ,start , err int  = 1 , 2 ,3
  two := &Node{
    Value : 2 ,
  }
  one.right = two
  open.Run("https://google.com/")
  open.Start("https://google.com")
  printNode(one)
}


func printNode( n *Node ) {
  fmt.Println("Value : ", n.Value)
  fmt.Println()
}
