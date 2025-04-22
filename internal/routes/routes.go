package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zero-remainder/go-ranker/internal/controllers"
)

func SetupPublicRoutes(app *fiber.App) {
	api := app.Group("/api")
	auth := api.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	seoController := controllers.NewSEOController()

	api.Get("/analyze", seoController.AnalyzeWebsite)
}
