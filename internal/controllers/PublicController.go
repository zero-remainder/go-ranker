package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type PublicController struct{}

func NewPublicController() *PublicController {
	return &PublicController{}
}

func (p *PublicController) Traffic(c *fiber.Ctx) error {
	//collection := database.GetCollection(constants.AnalyzeRequestsCollection)
	//
	//cursor, err := collection.Find(context.TODO(), bson.M{})
	//if err != nil {
	//	log.Println("Error fetching records:", err)
	//	return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch records"})
	//}
	//defer cursor.Close(context.TODO())
	//
	//var records []bson.M
	//if err := cursor.All(context.TODO(), &records); err != nil {
	//	log.Println("Error decoding records:", err)
	//	return c.Status(500).JSON(fiber.Map{"error": "Failed to decode records"})
	//}
	//
	//return c.JSON(fiber.Map{"message": "Records fetched successfully", "data": records})
	return c.JSON(fiber.Map{"message": "Records fetched successfully", "data": "test"})
}

func (p *PublicController) Analyze(c *fiber.Ctx) error {
	//url := c.Query("url")
	//if err := utils.ValidateURL(url); err != nil {
	//	return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	//}
	//collection := database.GetCollection(constants.AnalyzeRequestsCollection)
	//
	//clientIP := c.IP()
	//userAgent := c.Get("User-Agent")
	//host := c.Hostname()
	//
	//data := models.AnalyzeRequest{
	//	URL:       url,
	//	IP:        clientIP,
	//	UserAgent: userAgent,
	//	Host:      host,
	//	CreatedAt: time.Now(),
	//	Status:    false,
	//}
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//_, err := collection.InsertOne(ctx, data)
	//if err != nil {
	//	return c.JSON(fiber.Map{"status": false, "message": "Error inserting test record"})
	//}
	//
	//return c.JSON(fiber.Map{"status": true, "message": "URL validated successfully", "url": url})
	return c.JSON(fiber.Map{"status": true, "message": "URL validated successfully", "url": "test"})
}
