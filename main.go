package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waleedElgazar/resturant/configration"
	"github.com/waleedElgazar/resturant/routes"
)

func main() {
	configration.OpenConnection()
	app := fiber.New()
	routes.UserSetUp(app)
	routes.OrderSetUp(app)
	routes.PaymentSetUp(app)
	app.Listen(":8000")
}