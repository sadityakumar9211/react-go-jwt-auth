package main

import (
	"github.com/gofiber/fiber"
	"github.com/sadityakumar9211/jwt-auth-go/database"
	"github.com/sadityakumar9211/jwt-auth-go/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// }))

	routes.Setup(app)

	app.Listen(":8000")
}
