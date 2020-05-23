package main

import (
  "fmt"
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

  two := &Node{
    Value : 2 ,
  }
  one.right = two


  printNode(one)
}


func printNode( n *Node ) {
  fmt.Println("Value : ", n.Value)
  fmt.Println()
}
