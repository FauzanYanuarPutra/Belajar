package main

import (
	"go-paslon/config"
	"go-paslon/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
		config.ConnectDB()
    routes.PaslonRoutes(r)
    r.Run(":5000")
}





