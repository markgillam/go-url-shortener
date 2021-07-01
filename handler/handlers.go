package handler

import (
	"net/http"

	"github.com/markgillam/go-url-shortener/shortener"
	"github.com/markgillam/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)
//No binding for user_id so that it is optional
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9888/"
	c.JSON(http.StatusOK, gin.H{
		"response": host + shortUrl,
		"status":   http.StatusOK})

}
//GET to redirect, currently not working. 
//Possibly due to Gin HTTP handler
func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
//GET to display
func HandleShortUrlExpand(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.JSON(http.StatusOK, gin.H{
		"response": initialUrl,
		"status":   http.StatusOK})

}
