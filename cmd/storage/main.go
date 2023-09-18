package main

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
)
func main() {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		url := c.Request.URL
		log.Println("storage:", url.String())
		c.Next()
	})
	router.StaticFS("/static", http.Dir("./static/"))
	router.Run(":8888")
}
