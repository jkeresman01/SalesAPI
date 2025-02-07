package Controller

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/jkeresman01/SalesAPI/Config"
	"github.com/jkeresman01/SalesAPI/Model"
	"strconv"
	"time"
)

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}

	if data["passcode"] == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cashier passcode is required",
		})
	}

	if data["name"] == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cashier name is required",
		})
	}

	cashier := Model.Cashier{
		Name:      data["name"],
		Passcode:  data["passcode"],
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	db.DB.Create(&cashier)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Cashier added successfully",
		"data":    cashier,
	})
}

func GetCashiers(c *fiber.Ctx) error {
	var cashiers []Model.Cashier

	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))

	var count int64

	db.DB.Select("*").Limit(limit).Offset(skip).Find(&cashiers).Count(&count)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Cashiers",
		"data":    cashiers,
	})
}

func GetCashierWithId(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")

	var cashier Model.Cashier

	db.DB.Select("id, name").Where("id =?", cashierId).First(&cashier)

	cashierData := make(map[string]interface{})

	if cashier.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
			"error":   cashierData,
		})
	}

	cashierData["cashierId"] = cashier.Id
	cashierData["name"] = cashier.Name
	cashierData["createAt"] = cashier.CreatedAt
	cashierData["updateAt"] = cashier.UpdatedAt

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    cashierData,
	})
}

func UpdateCashierWithId(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")

	var cashier Model.Cashier

	db.DB.Select("*").Where("id=?", cashierId).Find(&cashier)

	if cashier.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
		})
	}

	var updatedCashier Model.Cashier
	err := c.BodyParser(&updatedCashier)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}

	if cashier.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Cashier name is requiered",
		})
	}

	db.DB.Save(&updatedCashier)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    updatedCashier,
	})
}

func DeleteCashierWithId(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")

	var cashier Model.Cashier

	db.DB.Where("id=?", cashierId).First(&cashier)

	if cashier.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
		})
	}

	db.DB.Where("id=?", cashierId).Delete(&cashier)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Cashier delete successfully",
	})
}
