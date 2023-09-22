package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		validated := false

		if !validated {
			// c.JSON(400, "the field is required")
			c.AbortWithStatusJSON(400, gin.H{
				"message": "the field name is required",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "You are logged in.",
		})
	})

	app.Run(":8000")

}