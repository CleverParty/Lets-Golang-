package main

import (
   "fmt"
)

type vertex struct { // A struct is a collection of fields.
  X int
  Y int
  Q ,Z int // the declaration can also be done like this.
}

func main() {

  i , j := 10,11

  point := &i

  fmt.Println(*point)
  *point = 13
  fmt.Println(i)

  point = &j
  fmt.Println(*point)
  *point = 14

  // struct testing :
  fmt.Println("Struct Testing")
  v := vertex{57, 3 }
  v.X = 63
  fmt.Println(v.X)
  // pointers can also be made to structs

  point2 := &v
  point2.X = 404
  fmt.Println("point2 referenced value : ",v.X)

}
