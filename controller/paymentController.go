package controller

import (
	"fmt"
	"time"

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
	currentTime := time.Now()
	payment:=models.Payment{
		PaymentDate: currentTime.Format("2006.01.02 15:04:05"),
		PaymentType: data.PaymentType,
		UserId: data.UserId,
		OrderId: data.OrderId,
		Amount: data.Amount,
	}
	err=database.AddPaymentForOrder(payment)
	if err!=nil {
		ctx.Status(fiber.ErrBadRequest.Code)
		return ctx.JSON(
			fiber.Map{
				"message":"order id or user id are not found",
			},
		)
	}
	return ctx.JSON(data)
}
