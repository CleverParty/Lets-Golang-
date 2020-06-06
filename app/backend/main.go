package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	a := [3]int{1, 2, 3} // first method
	var b [3]string      // second method of declaring the arrays
	fmt.Println(a)
	var avg, sumVal, storeDif int
	var y, j int
	a[0] = rand.Intn(100)
	a[1] = rand.Intn(100)
	a[2] = rand.Intn(100)
	b[0] = "India"
	b[1] = "Canada"
	b[2] = "Japan"
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		dst := "/Users/shanmukhasurapuraju/Documents"
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":3003")
}
