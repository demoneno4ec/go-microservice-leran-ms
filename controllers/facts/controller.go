package facts

import (
	"github.com/demoneno4ec/go-microservice-leran-ms/common/response"
	"github.com/demoneno4ec/go-microservice-leran-ms/database"
	"github.com/demoneno4ec/go-microservice-leran-ms/models"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Недостаточно данных")
	}

	database.DB.Db.Create(&fact)

	return response.EmptySuccess(c)
}

func ListFacts(c *fiber.Ctx) error {
	var facts []models.Fact
	database.DB.Db.Find(&facts)

	return response.Success(c, facts)
}

func Show(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).First(&fact)
	if result.Error != nil {
		return response.Error(c, fiber.StatusNotFound, "Факт не найден")
	}

	return response.Success(c, fact)
}

func Update(c *fiber.Ctx) error {
	fact := new(models.Fact)
	id := c.Params("id")

	// Parsing the request body
	if err := c.BodyParser(fact); err != nil {
		return response.Error(c, fiber.StatusServiceUnavailable, "Недостаточно данных")
	}

	result := database.DB.Db.Where("id = ?", id).First(&fact)
	if result.Error != nil {
		return response.Error(c, fiber.StatusNotFound, "Факт не найден")
	}

	result.Updates(fact)

	return response.EmptySuccess(c)
}

func Delete(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).First(&fact)
	if result.Error != nil {
		return response.Error(c, fiber.StatusNotFound, "Факт не найден")
	}

	database.DB.Db.Where("id = ?", id).Delete(&fact)

	return response.EmptySuccess(c)
}
