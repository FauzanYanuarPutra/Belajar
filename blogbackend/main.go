package main

import (
	"blogbackend/database"
	"blogbackend/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.ErrEnv()
	port:=os.Getenv("PORT")
	app:=fiber.New()
	routes.Setup(app)
	app.Listen(":"+port)

}
