package facts

import (
	"fmt"
	"github.com/demoneno4ec/go-microservice-leran-ms/database"
	"github.com/demoneno4ec/go-microservice-leran-ms/models"
	"github.com/gofiber/fiber/v2"
)

// todo вынести это в common/response
type errorResponse struct {
	Message string `json:"message"`
}

func Create(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(fiber.StatusOK).JSON(fact)
}

func ListFacts(c *fiber.Ctx) error {
	var facts []models.Fact
	database.DB.Db.Find(&facts)

	return c.Status(fiber.StatusOK).JSON(facts)
}

func Show(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).First(&fact)
	if result.Error != nil {
		errorResponse := errorResponse{}

		fmt.Println("test log 3")
		errorResponse.Message = "test 3"
		return c.Status(fiber.StatusNotFound).JSON(errorResponse)
	}

	return c.Status(fiber.StatusOK).JSON(fact)
}

func Update(c *fiber.Ctx) error {
	fact := new(models.Fact)
	id := c.Params("id")

	// Parsing the request body
	if err := c.BodyParser(fact); err != nil {
		errorResponse := errorResponse{}

		errorResponse.Message = "test 2"
		return c.Status(fiber.StatusServiceUnavailable).JSON(errorResponse)
	}

	result := database.DB.Db.Where("id = ?", id).First(&fact)
	if result.Error != nil {
		errorResponse := errorResponse{}

		fmt.Println("test log 3")
		errorResponse.Message = "test 3"
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "not updated",
		})
	}

	result.Updates(fact)

	return c.Status(fiber.StatusOK).JSON(fact)
}

func NotFound(c *fiber.Ctx) error {
	errorResponse := errorResponse{}

	errorResponse.Message = "test"
	return c.Status(fiber.StatusNotFound).JSON(errorResponse)
}
