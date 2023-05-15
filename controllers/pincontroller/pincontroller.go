package pincontroller

import (
	"brm_api/models"
	"brm_api/services/db/postgres"
	"brm_api/utils/common"
	"brm_api/utils/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func List(c *fiber.Ctx) error {
	resType := []map[string]interface{}{}
	var res *gorm.DB
	db := postgres.Connection
	packageID := c.QueryInt("package_id")

	user := c.Locals(models.UserKey).(*models.User)
	//user, _ := userservice.GetUserByID(4)
	sqlStr := "SELECT pins.id, pins.pin, pins.status, pin_transactions.id AS transaction_id, pin_transactions.shared_to AS shared_to FROM pins INNER JOIN pin_transactions AS pin_transactions ON pins.id = pin_transactions.pin_id where pins.status = ? and pins.package_id = ? and pin_transactions.shared_to = ?"
	if user.Role == models.SuperAdmin || user.Role == models.Admin {
		res = db.Raw(sqlStr, "un_used", packageID, 0).Scan(&resType)
	} else {
		res = db.Raw(sqlStr, "un_used", packageID, user.ID).Scan(&resType)
	}

	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.HTTPResponse{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	return c.JSON(response.HTTPResponse{
		Success: true,
		Message: "got all pins",
		Data:    resType,
	})

}

func Create(c *fiber.Ctx) error {
	user := c.Locals(models.UserKey).(*models.User)
	if user.Role == models.SuperAdmin || user.Role == models.Admin {
		reqType := new(createReqBody)
		if err := c.BodyParser(reqType); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.HTTPResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
		}
		errors := response.ValidateStruct(*reqType)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.HTTPResponse{
				Success: false,
				Message: "Field validation failed",
				Data:    errors,
			})
		}

		var pins []models.Pin
		db := postgres.Connection
		err := db.Transaction(func(tx *gorm.DB) error {
			for i := 0; i < reqType.Size; i++ {
				pin := models.Pin{
					Status:      models.UnUsedStatus,
					GeneratedBy: user.ID,
					PackageID:   reqType.PackageID,
					Pin:         common.GenerateRandString(5),
				}
				if err := tx.Create(&pin).Error; err != nil {
					return err
				}

				pinT := models.PinTransaction{
					PinID:    pin.ID,
					SharedTo: 0, //reqType.ShareTo,//todo: enable if want to generate to pins to specific user
				}
				if err := tx.Create(&pinT).Error; err != nil {
					return err
				}
				pins = append(pins, pin)
			}
			return nil
		})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.HTTPResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
		}
		return c.JSON(response.HTTPResponse{
			Success: true,
			Message: "pins created",
			Data:    pins,
		})
	} else {
		return c.Status(fiber.StatusForbidden).JSON(response.HTTPResponse{
			Success: false,
			Message: "Dont have enough permission to create Pins",
			Data:    nil,
		})
	}

}

func TransferPins(c *fiber.Ctx) error {
	user := c.Locals(models.UserKey).(*models.User)
	if user.Role == models.SuperAdmin || user.Role == models.Admin {
		reqType := new(transferReqBody)
		if err := c.BodyParser(reqType); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.HTTPResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
		}
		errors := response.ValidateStruct(*reqType)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.HTTPResponse{
				Success: false,
				Message: "Field validation failed",
				Data:    errors,
			})
		}

		db := postgres.Connection
		err := db.Transaction(func(tx *gorm.DB) error {
			for _, pinTransID := range reqType.TransactionID {
				tx.Model(&models.PinTransaction{}).Where("id = ?", pinTransID).Update("shared_to", reqType.ShareTo)
			}
			return nil
		})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.HTTPResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
		}
		return c.JSON(response.HTTPResponse{
			Success: true,
			Message: "pins transferred",
			Data:    nil,
		})
	} else {
		return c.Status(fiber.StatusForbidden).JSON(response.HTTPResponse{
			Success: false,
			Message: "Dont have enough permission to create Pins",
			Data:    nil,
		})
	}

}

// func Create(c *fiber.Ctx) error {
// 	user := c.Locals(models.UserKey).(*models.User)
// 	if user.Role == models.SuperAdmin || user.Role == models.Admin {
// 		reqType := new(createReqBody)
// 		if err := c.BodyParser(reqType); err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(response.HTTPResponse{
// 				Success: false,
// 				Message: err.Error(),
// 				Data:    nil,
// 			})
// 		}
// 		errors := response.ValidateStruct(*reqType)
// 		if errors != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(response.HTTPResponse{
// 				Success: false,
// 				Message: "Field validation failed",
// 				Data:    errors,
// 			})
// 		}

// 		var pins []models.Pin
// 		for i := 0; i < reqType.Size; i++ {
// 			pin := models.Pin{
// 				Status:      models.UnUsedStatus,
// 				GeneratedBy: user.ID,
// 				ShareTo:     reqType.ShareTo,
// 				Pin:         common.GenerateRandString(5),
// 			}
// 			pins = append(pins, pin)
// 		}
// 		db := postgres.Connection
// 		res := db.Create(&pins)
// 		if res.Error != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(response.HTTPResponse{
// 				Success: false,
// 				Message: res.Error.Error(),
// 				Data:    nil,
// 			})
// 		}
// 		return c.JSON(response.HTTPResponse{
// 			Success: true,
// 			Message: "pins created",
// 			Data:    pins,
// 		})
// 	} else {
// 		return c.Status(fiber.StatusForbidden).JSON(response.HTTPResponse{
// 			Success: false,
// 			Message: "Dont have enough permission to create Pins",
// 			Data:    nil,
// 		})
// 	}

// }

func Update(c *fiber.Ctx) error {

	return nil

}

func Delete(c *fiber.Ctx) error {

	return nil

}
