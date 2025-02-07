package Controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	db "github.com/jkeresman01/SalesAPI/Config"
	"github.com/jkeresman01/SalesAPI/Midleware"
	"github.com/jkeresman01/SalesAPI/Model"
	"strconv"
)

func CreateCategory(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		log.Fatalf("Category error in post request %v", err)
	}

	if data["name"] == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Category name is required",
			"error":   map[string]interface{}{},
		})
	}

	category := Model.Category{
		Name: data["name"],
	}

	db.DB.Create(&category)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    category,
	})
}

func GetCategoryWithId(c *fiber.Ctx) error {
	categoryId := c.Params("categoryId")

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

	var category Model.Category
	db.DB.Select("id, name").Where("id=?", categoryId).First(&category)

	categoryData := make(map[string]interface{})
	categoryData["categoryId"] = category.Id
	categoryData["name"] = category.Name

	if category.Name == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "No category found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    categoryData,
	})
}

func DeleteCategoryWithId(c *fiber.Ctx) error {
	categoryId := c.Params("categoryId")

	var category Model.Category
	db.DB.Where("id=?", categoryId).First(&category)

	if category.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": true,
			"message": "Category not found!",
			"error":   map[string]interface{}{},
		})
	}

	db.DB.Where("id=?", categoryId).Delete(&category)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
	})
}

func UpdateCategoryWithId(c *fiber.Ctx) error {
	categoryId := c.Params("categoryId")
	var category Model.Category

	db.DB.Find(&category, "id=?", categoryId)

	if category.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Category doesn't exist!",
		})
	}

	var updatedCategoryData Model.Category
	_ = c.BodyParser(&updatedCategoryData)

	if updatedCategoryData.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Category name is required!",
			"error":   map[string]interface{}{},
		})
	}

	category.Name = updatedCategoryData.Name

	db.DB.Save(&category)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    category,
	})
}

func GetCategories(c *fiber.Ctx) error {
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

	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))

	var count int64
	var category []Model.Category

	db.DB.Select("id, name").Limit(limit).Offset(skip).Find(&category).Count(&count)

	metaMap := map[string]interface{}{
		"total": count,
		"limit": limit,
		"skip":  skip,
	}

	categoriesData := map[string]interface{}{
		"categories": category,
		"meta":       metaMap,
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    categoriesData,
	})
}
