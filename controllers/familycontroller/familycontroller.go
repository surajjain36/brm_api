package familycontroller

import (
	"brm_api/models"
	"brm_api/services/db/postgres"
	"brm_api/utils/response"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func Create(c *fiber.Ctx) error {

	return nil

}

func Update(c *fiber.Ctx) error {

	return nil

}

func Delete(c *fiber.Ctx) error {

	return nil

}

func List(c *fiber.Ctx) error {

	resType := []models.Family{}

	db := postgres.Connection
	res := db.Preload(clause.Associations).Find(&resType)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.HTTPResponse{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	return c.JSON(response.HTTPResponse{
		Success: true,
		Message: "Success get all Users.",
		Data:    resType,
	})

}

func First(c *fiber.Ctx) error {

	resType := models.Family{}
	db := postgres.Connection
	id := c.Params("id")
	idUUid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.HTTPResponse{
			Success: false,
			Message: "id is not a valid uuid",
			Data:    nil,
		})
	}
	res := db.Preload(clause.Associations).First(&resType, &models.Family{
		ID: idUUid,
	})
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.HTTPResponse{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	return c.JSON(response.HTTPResponse{
		Success: true,
		Message: "got an individual",
		Data:    resType,
	})

}
