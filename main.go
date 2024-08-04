package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/generate", func(c *gin.Context) {
		random := rand.Int()
		c.String(http.StatusOK, "Random Value is %v\n", random)
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome to Go gin server")
	})
	router.GET("/admin", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to admin panel")
	})

	router.Run(":8080")
}
