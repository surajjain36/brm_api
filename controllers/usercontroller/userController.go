package usercontroller

import (
	"brm_api/controllers/pincontroller"
	"brm_api/models"
	"brm_api/services/db/postgres"
	"brm_api/utils/response"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func List(c *fiber.Ctx) error {
	user := c.Locals(models.UserKey).(*models.User)
	resType := []models.User{}
	db := postgres.Connection
	var res *gorm.DB
	perPage, _ := strconv.Atoi(c.Query("per_page", strconv.Itoa(10)))
	page, _ := strconv.Atoi(c.Query("page", "1"))

	where := ""
	uIDStr := strconv.FormatUint(uint64(user.ID), 10)
	if user.Role == models.SuperAdmin || user.Role == models.Admin {
		where = fmt.Sprintf(`id NOT IN (%s) AND partner_id is distinct from %s AND referer_id != 0`, uIDStr, uIDStr)
		// } else if user.Role == models.Manager || user.Role == models.Leader {
		// 	where = fmt.Sprintf(`id NOT IN (%s) AND partner_id is distinct from %s AND hierarchy LIKE '%%|%s|%%' AND referer_id != 0`, uIDStr, uIDStr, uIDStr)
	} else if user.Role == models.Customer {
		where = fmt.Sprintf(`id NOT IN (%s) AND partner_id is distinct from %s AND referer_ id = %s AND referer_id != 0`, uIDStr, uIDStr, uIDStr)
	}
	res = db.Debug().Preload(clause.Associations).Where(where).Limit(perPage).Offset(page - 1).Find(&resType)

	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.HTTPResponse{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	return c.JSON(response.HTTPResponse{
		Success: true,
		Message: "got all users",
		Data:    resType,
	})

}

func First(c *fiber.Ctx) error {
	resType := models.User{}
	db := postgres.Connection
	id := c.Params("id")
	uintID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.HTTPResponse{
			Success: false,
			Message: "id is not a valid int id",
			Data:    nil,
		})
	}
	res := db.Preload(clause.Associations).First(&resType, &models.User{
		ID: uint(uintID),
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
		Message: "got an user",
		Data:    resType,
	})
}

func Create(c *fiber.Ctx) (err error) {
	reqPayload := new(models.User)
	if err := c.BodyParser(reqPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.HTTPResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	errors := response.ValidateStruct(*reqPayload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.HTTPResponse{
			Success: false,
			Message: "Field validation failed",
			Data:    errors,
		})
	}

	user := c.Locals(models.UserKey).(*models.User)
	db := postgres.Connection
	reqPayload.Hierarchy = user.Hierarchy + "|" + strconv.FormatUint(uint64(user.ID), 10) + "|"
	reqPayload.CreatedBy = user.ID
	if reqPayload.RefererID == 0 {
		reqPayload.PartnerID = user.ID
	}

	res := db.Create(&reqPayload)

	for _, pkg := range reqPayload.Packages {
		newPin := models.Pin{
			ID:         pkg.PinID,
			Status:     models.UsedStatus,
			UsedBy:     user.ID,
			AssignedTo: reqPayload.ID,
		}
		err = pincontroller.UpdatePin(newPin)
	}

	if res.Error != nil || err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.HTTPResponse{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	return c.JSON(response.HTTPResponse{
		Success: true,
		Message: "created user",
	})
}

func Roles(c *fiber.Ctx) error {
	var res *gorm.DB
	resType := []models.Role{}
	db := postgres.Connection

	user := c.Locals(models.UserKey).(*models.User)
	if user.Role == models.SuperAdmin {
		res = db.Preload(clause.Associations).Find(&resType)
	} else {
		res = db.Preload(clause.Associations).Where("key NOT IN (?,?)", models.SuperAdmin, models.Admin).Find(&resType)
	}
	// else if user.Role == models.Admin {
	// 	res = db.Preload(clause.Associations).Where("key NOT IN (?,?)", models.SuperAdmin, models.Admin).Find(&resType)
	// }
	// else if user.Role == models.Manager {
	// 	res = db.Preload(clause.Associations).Where("key NOT IN (?,?,?)", models.SuperAdmin, models.Admin, models.Manager).Find(&resType)
	// } else if user.Role == models.Leader {
	// 	res = db.Preload(clause.Associations).Where("key NOT IN (?,?,?,?)", models.SuperAdmin, models.Admin, models.Manager, models.Leader).Find(&resType)
	// }

	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.HTTPResponse{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	return c.JSON(response.HTTPResponse{
		Success: true,
		Message: "got all roles",
		Data:    resType,
	})

}

func Packages(c *fiber.Ctx) error {
	resType := []models.Package{}
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
		Message: "got all roles",
		Data:    resType,
	})

}

func Update(c *fiber.Ctx) error {
	updatePayload := new(models.User)
	if err := c.BodyParser(updatePayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.HTTPResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	errors := response.ValidateStruct(*updatePayload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.HTTPResponse{
			Success: false,
			Message: "Field validation failed",
			Data:    errors,
		})
	}

	db := postgres.Connection
	res := db.Preload(clause.Associations).Model(&updatePayload).Updates(updatePayload)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.HTTPResponse{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	return c.JSON(response.HTTPResponse{
		Success: true,
		Message: "user updated",
	})

}

func Delete(c *fiber.Ctx) error {

	return nil

}
