package app

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
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
