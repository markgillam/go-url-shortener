package main

import (
	"fmt"
	"net/http"
	"github.com/markgillam/go-url-shortener/handler"
	"github.com/markgillam/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)
// Root URL welcome to be tested in browser
func main() {
	gin.SetMode(gin.ReleaseMode) //comment out this line to return to debug mode
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the URL Shortener API", "status": http.StatusOK})

	})

	// Can't get redirect working, possibly due to using Gin
	//r.GET("/:shortUrl", func(c *gin.Context) {
	//	handler.HandleShortUrlRedirect(c)
	//})
	
	//Shorten POST
	r.POST("/api/v1/shorten/", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	//Expand GET
	r.GET("/api/v1/lookup/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlExpand(c)
	})

	//Redis store initialize and port selection
	store.InitializeStore()

	err := r.Run(":9888")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
