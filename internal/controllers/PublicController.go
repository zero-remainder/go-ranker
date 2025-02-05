package controllers

import "github.com/gofiber/fiber/v2"

type PublicController struct{}

func NewPublicController() *PublicController {
	return &PublicController{}
}

func (pc *PublicController) GetPublic(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": true, "message": "Go Ranker"})
}
