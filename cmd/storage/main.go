package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	router := gin.Default()
	router.StaticFS("/static", http.Dir("./static/"))
	router.Run()
}
