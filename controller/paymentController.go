package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/waleedElgazar/resturant/database"
	"github.com/waleedElgazar/resturant/models"
)

func PayForOrder(ctx *fiber.Ctx) error {
	var data models.Payment
	err := ctx.BodyParser(&data)
	if err != nil {
		fmt.Println("error while parsing data", err)
		return err
	}
	err=database.AddPaymentForOrder(data)
	if err!=nil {
		ctx.Status(fiber.ErrBadRequest.Code)
		return ctx.JSON(
			fiber.Map{
				"message":"orderid or user id are not found",
			},
		)
	}
	return ctx.JSON(data)
}
