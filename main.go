package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Set the router
	router = gin.Default()

	router.LoadHTMLGlob("templates/")

}
