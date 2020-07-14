package main

import (
	"fmt"
	"image/png"
	"os"

	"image"

	"github.com/skratchdot/open-golang/open"
)

type Node struct {
	Value, payload int
	left           *Node
	right          *Node
}

func main() {

	myImage := image.NewRGBA(image.Rect(0, 0, 10, 4)) // imgae generation

	myImage.Pix[0] = 255 // 1st pixel red
	myImage.Pix[1] = 0   // 1st pixel green
	outputFile, err := os.Create("works" +
		".png") // storeing encoded image in a .png file
	if err != nil {
		fmt.Println("error occured")
	}
	png.Encode(outputFile, myImage)

	outputFile.Close()

	one := &Node{
		Value:   10,
		payload: 0,
	}
	// var val ,start , err int  = 1 , 2 ,3
	two := &Node{
		Value:   2,
		payload: 0,
	}
	one.right = two
	open.Run("https://google.com/")
	// open.Start("https://google.com")
	printNode(one)

}

func access(n *Node) {
	fmt.Println("Value : ", n.payload)
	fmt.Println()
}

func printNode(n *Node) {
	fmt.Println("Value : ", n.Value)
	fmt.Println()
}
