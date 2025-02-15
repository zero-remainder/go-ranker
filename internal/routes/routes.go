package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zero-remainder/go-ranker/internal/controllers"
)

func SetupPublicRoutes(app *fiber.App) {
	publicController := controllers.NewPublicController()
	apiGroup := app.Group("/api")

	apiGroup.Get("/analyze", publicController.Analyze)
	apiGroup.Get("/traffic-records", publicController.Traffic)
}
