package packagecontroller

import (
	"brm_api/models"
	"brm_api/services/db/postgres"
	"brm_api/utils/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ListAchievers(c *fiber.Ctx) error {
	resType := []models.User{}
	var res *gorm.DB
	db := postgres.Connection
	perPage, _ := strconv.Atoi(c.Query("per_page", strconv.Itoa(10)))
	page, _ := strconv.Atoi(c.Query("page", "1"))

	user := c.Locals(models.UserKey).(*models.User)
	//uIDStr := strconv.FormatUint(uint64(user.ID), 10)
	joinQuery := "JOIN user_packages AS up ON up.user_id=users.id"
	if user.Role == models.SuperAdmin || user.Role == models.Admin {
		res = db.Debug().Preload(clause.Associations).Joins(joinQuery).Where("up.package = ? AND referer_id != 0", models.AchieverKey).Limit(perPage).Offset(page - 1).Find(&resType)
		//res = db.Debug().Preload("Packages", "package = ?", "achevers").Find(&resType, "users.referer_id = ?", user.ID)
		// } else if user.Role == models.Manager || user.Role == models.Leader {
		// 	where := fmt.Sprintf(`up.package = '%s' AND hierarchy LIKE '%%|%s|%%' AND referer_id != 0`, models.AchieverKey, uIDStr)
		// 	res = db.Debug().Joins(joinQuery).Where(where).Limit(perPage).Offset(page-1).Find(&resType)
	} else if user.Role == models.Customer {
		res = db.Debug().Joins(joinQuery).Where("up.package = ? AND referer_id != 0", models.AchieverKey).Limit(perPage).Offset(page-1).Find(&resType, "users.referer_id = ?", user.ID)
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
		Message: "got all achievers",
		Data:    resType,
	})

}

func ListClassicUsers(c *fiber.Ctx) error {
	resType := []models.User{}
	var res *gorm.DB
	db := postgres.Connection
	perPage, _ := strconv.Atoi(c.Query("per_page", strconv.Itoa(10)))
	page, _ := strconv.Atoi(c.Query("page", "1"))

	user := c.Locals(models.UserKey).(*models.User)
	//uIDStr := strconv.FormatUint(uint64(user.ID), 10)
	joinQuery := "JOIN user_packages AS up ON up.user_id=users.id"
	if user.Role == models.SuperAdmin || user.Role == models.Admin {
		res = db.Debug().Preload(clause.Associations).Joins(joinQuery).Where("up.package = ? AND referer_id != 0", models.ClassicKey).Limit(perPage).Offset(page - 1).Find(&resType)
		//res = db.Debug().Preload("Packages", "package = ?", "achevers").Find(&resType, "users.referer_id = ?", user.ID)
		// } else if user.Role == models.Manager || user.Role == models.Leader {
		// 	where := fmt.Sprintf(`up.package = '%s' AND hierarchy LIKE '%%|%s|%%' AND referer_id != 0`, models.ClassicKey, uIDStr)
		// 	res = db.Debug().Joins(joinQuery).Where(where).Limit(perPage).Offset(page - 1).Find(&resType)
	} else if user.Role == models.Customer {
		res = db.Debug().Joins(joinQuery).Where("up.package = ? AND referer_id != 0", models.ClassicKey).Limit(perPage).Offset(page-1).Find(&resType, "users.referer_id = ?", user.ID)
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
		Message: "got all achievers",
		Data:    resType,
	})

}

func ListClassic(c *fiber.Ctx) error {
	resType := []models.User{}

	return c.JSON(response.HTTPResponse{
		Success: true,
		Message: "got all classic users",
		Data:    resType,
	})

}
