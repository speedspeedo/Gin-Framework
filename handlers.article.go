package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(ctx *gin.Context) {
	articles := getAllArticles()

	// Call the render function with the name of the template to render
	render(ctx, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
}

func getArticle(ctx *gin.Context) {
	// Check if the ID is article ID is valid
	if articleID, err := strconv.Atoi(ctx.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// Call the HTML method of the Context to render a template
			ctx.HTML(
				// Set the HTTP status to 200 (OK)
				http.StatusOK,
				// Use the index.html template
				"article.html",
				// Pass the data that the page uses
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)
		} else {
			// If the article is not found, abort with an error
			ctx.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		ctx.AbortWithError(http.StatusNotFound, err)
	}
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that the template name is present
func render(ctx *gin.Context, data gin.H, templateName string) {
	switch ctx.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		fmt.Println(ctx, "000")
		ctx.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		ctx.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		fmt.Println(ctx, "111")
		ctx.HTML(http.StatusOK, templateName, data)
	}
}
