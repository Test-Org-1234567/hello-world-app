#test
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//create a router and corresponding groups
	router := gin.New()
	router.StaticFS("/", http.Dir("templates"))

	fmt.Println("running")

	err := router.Run(":9090")
	if err != nil {
		panic(err)
	}
}
