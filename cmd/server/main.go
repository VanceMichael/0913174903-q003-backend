package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func router() http.Handler {
	r := gin.New()
	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
	return r
}

func main() {
	port := os.Getenv("PORT")
	if port == "" { port = "8080" }
	_ = http.ListenAndServe(":"+port, router())
}
