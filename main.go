package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const port = ":3000"

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola, mundo!",
		})
	})

	router.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola, mundo!",
		})
	})

	router.POST("/team/pokemons", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola, mundo!",
		})
	})

	router.GET("/team", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola, mundo!",
		})
	})

	router.DELETE("/team/pokemons/:pokeid", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola, mundo!",
		})
	})

	router.PUT("/team", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola, mundo!",
		})
	})
	return router
}

func main() {
	router := setupRouter()
	fmt.Printf("Server started at port %s\n", port)
	router.Run(port)
}
