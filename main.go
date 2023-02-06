package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Set the router
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.
	router.GET("/", func(ctx *gin.Context) {

		// Call the HTML method of the Context to render a template
		ctx.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'titel')
			gin.H{
				"title": "Home page",
			},
		)
	})

	router.Run()
}
