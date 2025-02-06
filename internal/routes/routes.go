package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zero-remainder/go-ranker/internal/controllers"
)

func SetupPublicRoutes(app *fiber.App) {
	publicController := controllers.NewPublicController()
	//userGroup := app.Group("/web")
	//userGroup.Get("/public", publicController.GetPublic)
	app.Get("/", publicController.Index)
	app.Get("/traffic", publicController.Traffic)
}
