package routes

import (
	"github.com/gofiber/fiber"
	"github.com/sadityakumar9211/jwt-auth-go/controllers"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

}
