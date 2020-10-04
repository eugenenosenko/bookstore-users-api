package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapURLs()
	err := router.Run(":8080")

	if err != nil {
		fmt.Println(err.Error())
	}
}
