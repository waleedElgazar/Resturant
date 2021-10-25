package models

type TypePayment string
const(
	Cash	TypePayment	="cash"
	Visa	="visa"
)

type Payment struct {
	Id			uint		`json:"-"`
	UserId		uint		`json:"user_id"`
	OrderId		uint		`json:"order_id"`
	PaymentDate string		`json:"data"`
	Amount		float64		`json:"total_price"`
	PaymentType	TypePayment	`json:"payment_type"`
}