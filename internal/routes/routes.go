package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zero-remainder/go-ranker/internal/controllers"
)

func SetupPublicRoutes(app *fiber.App) {
	publicController := controllers.NewPublicController()
	//userGroup.Get("/public", publicController.GetPublic)
	app.Get("/", publicController.Index)
	app.Get("/traffic", publicController.Traffic)

	apiGroup := app.Group("/api")

	apiGroup.Get("/analyze", publicController.Analyze)
}
