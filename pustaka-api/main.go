package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
		router := gin.Default()

		router.GET("/", rootHandler)
		router.GET("/about", aboutHandler)
		router.GET("/books/:id/:title", booksHandler)
		router.GET("/params", paramsHandler)
		router.POST("/books", postBooksHandler)

		router.Run(":3000")
}

func rootHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"name": "Muhmmad Fauzan Yanuar Putra",
	})
}

func aboutHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"about": "im fullstack developer & UI / UX DESIGNER",
	})
}

func booksHandler(c *gin.Context){
	id := c.Param("id")
	title := c.Param("title")


	c.JSON(http.StatusOK, gin.H{
		"id": id,
		"title": title,
	})
}


func paramsHandler(c *gin.Context){
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

type Book struct {
	Title string `json:"title" binding:"required"`
	Price int `json:"price" binding:"required,number"`
}

func postBooksHandler(c *gin.Context){
	var BookInput Book

	err := c.ShouldBindJSON(&BookInput)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H {
		// 	"error": err.Error(),
		// 	"message": "Error keep calm",
		// })
		c.JSON(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": BookInput.Title,
		"price": BookInput.Price,

	})


}