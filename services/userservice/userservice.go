package userservice

import (
	"brm_api/models"
	"brm_api/services/db/postgres"
	"brm_api/utils/apperror"

	"gorm.io/gorm/clause"
)

// GetUserByID from database
func GetUserByID(id uint) (models.User, error) {
	db := postgres.Connection
	user := models.User{
		ID: id,
	}
	res := db.First(&user)
	if res.Error != nil {
		return models.User{}, &apperror.AppError{
			Status: 500,
			Msg:    res.Error.Error(),
			Errors: nil,
		}
	}
	return user, &apperror.AppError{}
}

// GetUserByEmailID from database
func GetUserByEmailID(mobile string) (models.User, error) {
	db := postgres.Connection
	user := models.User{}
	res := db.First(&user, &models.User{
		EmailID: mobile,
	})
	if res.Error != nil {
		status := 500
		if res.Error.Error() == "record not found" {
			status = 404
		}
		return models.User{}, &apperror.AppError{
			Status: status,
			Msg:    res.Error.Error(),
			Errors: nil,
		}
	}
	return user, &apperror.AppError{}

}

// GetUserByMobileNumber from database
func GetUserByMobileNumber(mobile string) (models.User, error) {
	db := postgres.Connection
	user := models.User{}
	res := db.First(&user, &models.User{
		MobileNumber: mobile,
	})
	if res.Error != nil {
		status := 500
		if res.Error.Error() == "record not found" {
			status = 404
		}
		return models.User{}, &apperror.AppError{
			Status: status,
			Msg:    res.Error.Error(),
			Errors: nil,
		}
	}
	return user, &apperror.AppError{}

}

// UpdateUser from database
func UpdateUser(user models.User) error {
	db := postgres.Connection
	res := db.Preload(clause.Associations).Model(&user).Updates(user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// CreateUser from database
func CreateUser(user models.User) {
	db := postgres.Connection
	db.Create(&user)
}
