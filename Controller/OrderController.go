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

	ProductOrders := make([]Model.ProductResponseOrder, 0)

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

		ProductOrders = append(ProductOrders,
			&Model.ProductResponseOrder{
				ProductId:        productOrder.Id,
				Price:            productOrder.Price,
				Name:             productOrder.Name,
				Discount:         discount,
				Qty:              v.Quantity,
				TotalNormalPrice: productOrder.Price,
				TotalFinalPrice:  totalFinalPrice,
			})
	}
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
			"products": ProductOrders,
		},
	})

}

func GetSubTotalOrders(c *fiber.Ctx) error {
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

	type product struct {
		ProductId int `json:"productId"`
		Quantity  int `json:"qty"`
	}

	body := struct {
		Products []product `json:"products"`
	}{}

	if err := c.BodyParser(&body.Products); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Empty body",
		})
	}

	ResponseOrders := make([]*Model.ProductResponseOrder, 0)

	var TotalInvoicePrice = struct {
		totalPrice int
	}{}

	for _, v := range body.Products {
		totalPrice := 0

		prods := Model.ProductOrder{}

		var discount Model.Discount
		db.DB.Table("products").Where("id=?", v.ProductId).First(&prods)
		db.DB.Where("id=?", prods.DiscountId).First(&discount)

		disc := 0

		if discount.Type == "PERCENT" {
			totalPrice = prods.Price * v.Quantity
			percentage := totalPrice * discount.Result / 100
			disc = totalPrice - percentage
			TotalInvoicePrice.totalPrice = TotalInvoicePrice.totalPrice + disc
		}

		if discount.Type == "BUY_N" {
			totalPrice = prods.Price * v.Quantity
			disc = totalPrice - discount.Result
			TotalInvoicePrice.totalPrice = TotalInvoicePrice.totalPrice + disc
		}

		ResponseOrders = append(ResponseOrders,
			&Model.ProductResponseOrder{
				ProductId:        prods.Id,
				Name:             prods.Name,
				Price:            prods.Price,
				Discount:         discount,
				Qty:              v.Quantity,
				TotalNormalPrice: v.Quantity,
				TotalFinalPrice:  disc,
			})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data": map[string]interface{}{
			"subTotal": TotalInvoicePrice.totalPrice,
			"products": ResponseOrders,
		},
	})
}

func GetOrderWithId(c *fiber.Ctx) error {
	headerToken := c.Get("Authorization")

	if headerToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Authorization",
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

	orderId := c.Params("orderId")

	var order Model.Order

	db.DB.Where("id=?", orderId).First(&order)

	if order.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Order does not exist!",
		})
	}

	//TODO

	return nil
}

func CheckOrder(c *fiber.Ctx) error {
	orderId := c.Params("orderId")

	var order Model.Order

	db.DB.Where("id=?", orderId).First(&order)

	if order.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Order does not exist!",
		})
	}

	if order.IsDownloaded == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "success",
			"data": map[string]interface{}{
				"isDownloaded": false,
			},
		})
	}

	if order.IsDownloaded == 1 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "success",
			"data": map[string]interface{}{
				"isDownloaded": true,
			},
		})
	}

	return nil
}

func GetOrders(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))

	var count int64
	var order []Model.Order

	db.DB.Select("*").Limit(limit).Offset(skip).Find(&order).Count(&count)

	type OrderList struct {
		OrderId        int               `json:"orderId"`
		CashierID      int               `json:"cashiersId"`
		PaymentTypesId int               `json:"paymentTypesId"`
		TotalPrice     int               `json:"totalPrice"`
		TotalPaid      int               `json:"totalPaid"`
		TotalReturn    int               `json:"totalReturn"`
		ReceiptId      string            `json:"receiptId"`
		CreatedAt      time.Time         `json:"createdAt"`
		Payments       Model.PaymentType `json:"payment_type"`
		Cashiers       Model.Cashier     `json:"cashier"`
	}

	OrderResponse := make([]*OrderList, 0)

	for _, v := range order {
		cashier := Model.Cashier{}
		db.DB.Where("id=?", v.CashierId).Find(&cashier)

		paymentType := Model.PaymentType{}
		db.DB.Where("id=?", v.PaymentTypeId).Find(&paymentType)

		OrderResponse = append(OrderResponse, &OrderList{
			OrderId:        v.Id,
			CashierID:      v.CashierId,
			PaymentTypesId: v.PaymentTypeId,
			TotalPrice:     v.TotalPrice,
			TotalPaid:      v.TotalPaid,
			TotalReturn:    v.TotalReturn,
			ReceiptId:      v.ReceiptId,
			CreatedAt:      v.CreatedAt,
			Payments:       paymentType,
			Cashiers:       cashier,
		})
	}

	return c.Status(404).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    OrderResponse,
		"meta": map[string]interface{}{
			"total": count,
			"limit": limit,
			"skip":  skip,
		},
	})
}

func DownloadOrder(c *fiber.Ctx) error {
	return nil
}
