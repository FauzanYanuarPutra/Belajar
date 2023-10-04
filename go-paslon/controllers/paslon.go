package controllers

import (
	"go-paslon/config"
	"go-paslon/dto"
	"go-paslon/models"
	"go-paslon/service"

	"github.com/gin-gonic/gin"
)

func GetPaslons(c *gin.Context) {
    db := config.DB
    var paslons []models.Paslon
    if err := db.Find(&paslons); err.Error != nil {
        c.JSON(404, gin.H{"error": "Paslons not found"})
        return
    }
    c.JSON(200, paslons)
}

func GetPaslonByID(c *gin.Context) {
    db := config.DB
    var paslon models.Paslon
    id := c.Param("id")
    if err := db.First(&paslon, id); err.Error != nil {
        c.JSON(404, gin.H{"error": "Paslon not found"})
        return
    }
    c.JSON(200, paslon)
}

func CreatePaslon(c *gin.Context) {
    db := config.DB

	name := c.PostForm("name")
	visi := c.PostForm("visi")

	formfile, _, err := c.Request.FormFile("image")
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	uploadUrl, err := service.NewMediaUpload().FileUpload(models.File{File: formfile})
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	paslonModel := models.Paslon{
        Name:  name,
        Visi:  visi,
        Image: uploadUrl,
	}

	db.Create(&paslonModel)
    c.JSON(200, paslonModel)
}


func UpdatePaslon(c *gin.Context) {
	db := config.DB
	var paslonDTO dto.UpdatePaslonDTO
	id := c.Param("id")

	var paslonModel models.Paslon
	if err := db.First(&paslonModel, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Paslon not found"})
        return
	}

	if err := c.ShouldBindJSON(&paslonDTO); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
	}

	db.Model(&paslonModel).Updates(paslonDTO)
	c.JSON(200, paslonModel)
}


func DeletePaslon(c *gin.Context) {
    db := config.DB
    id := c.Param("id")
    db.Delete(&models.Paslon{}, id)

    c.JSON(200, gin.H{"message": "Paslon deleted"})
}

