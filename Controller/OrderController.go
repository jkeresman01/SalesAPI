package Controller

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/jkeresman01/SalesAPI/Config"
	"github.com/jkeresman01/SalesAPI/Midleware"
	"github.com/jkeresman01/SalesAPI/Model"
	"math/rand"
	"strconv"
	"time"
)

func CreateOrder(c *fiber.Ctx) error {
	headerToken := c.Get("Authorization")

	if headerToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"error":   map[string]interface{}{},
		})
	}

	token := Midleware.SplitToken(headerToken)

	if err := Midleware.AuthenticateToken(token); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"error":   map[string]interface{}{},
		})
	}

	type products struct {
		ProductId int `json:"productId"`
		Quantity  int `json:"qty"`
	}

	body := struct {
		PaymentId int        `json:"paymentId"`
		TotalPaid int        `json:"totalPaid"`
		Products  []products `json:"products"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Empty body",
			"error":   map[string]interface{}{},
		})
	}

	ProductResponse := make([]Model.ProductResponseOrder, 0)

	var TotalInvoicePrice = struct {
		totalPrice int
	}{}

	productIds := ""
	quantities := ""

	for _, v := range body.Products {
		totalPrice := 0
		productIds = productIds + "," + strconv.Itoa(v.ProductId)
		quantities = quantities + "," + strconv.Itoa(v.Quantity)

		productOrder := Model.ProductOrder{}
		var discount Model.Discount

		db.DB.Table("products").Where("id=?", &v.ProductId).First(&productOrder)
		db.DB.Where("id=?", productOrder.DiscountId).Find(&discount)

		totalFinalPrice := 0

		if discount.Type == "BUY_N" {
			totalPrice = productOrder.Price * v.Quantity

			totalFinalPrice = totalPrice - discount.Result
			TotalInvoicePrice.totalPrice = TotalInvoicePrice.totalPrice + totalFinalPrice
		}

		if discount.Type == "PERCENT" {
			totalPrice = productOrder.Price * v.Quantity
			percent := totalPrice * discount.Result / 100
			totalFinalPrice = totalPrice - percent
			TotalInvoicePrice.totalPrice = TotalInvoicePrice.totalPrice + totalFinalPrice
		}

		ProductResponse = append(ProductResponse,
			&Model.ProductResponseOrder{
				ProductId:        productOrder.Id,
				Name:             productOrder.Name,
				Price:            productOrder.Price,
				Discount:         discount,
				Qty:              v.Quantity,
				TotalNormalPrice: productOrder.Price,
				TotalFinalPrice:  totalFinalPrice,
			},
		)

		orderResponse := Model.Order{
			CashierId:     1,
			PaymentTypeId: body.PaymentId,
			TotalPrice:    TotalInvoicePrice.totalPrice,
			TotalPaid:     body.TotalPaid,
			TotalReturn:   body.TotalPaid - TotalInvoicePrice.totalPrice,
			ReceiptId:     "R000" + strconv.Itoa(rand.Intn(1000)),
			ProductId:     productIds,
			Quantities:    quantities,
			UpdatedAt:     time.Now().UTC(),
			CreatedAt:     time.Now().UTC(),
		}

		db.DB.Create(&orderResponse)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "success",
			"success": true,
			"data": map[string]interface{}{
				"order":    orderResponse,
				"products": ProductResponse,
			},
		})

	}
}

func GetOrders(c *fiber.Ctx) error {
	return nil
}

func GetOrderWithId(c *fiber.Ctx) error {
	return nil
}

func SubTotalOrders(c *fiber.Ctx) error {
	return nil
}

func DownloadOrder(c *fiber.Ctx) error {
	return nil
}

func CheckOrder(c *fiber.Ctx) error {
	return nil
}
