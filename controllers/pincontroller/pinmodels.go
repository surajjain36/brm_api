package pincontroller

import (
	"brm_api/models"
	"brm_api/services/db/postgres"
)

type createReqBody struct {
	Size      int  `json:"size"`
	ShareTo   uint `json:"share_to"`
	PackageID uint `json:"package_id"`
}

type transferReqBody struct {
	ShareTo       uint   `json:"share_to"`
	TransactionID []uint `json:"transaction_ids"`
}

// UpdatePin in database
func UpdatePin(pin models.Pin) error {
	db := postgres.Connection
	db.Model(&pin).
		Select("status", "assigned_to", "used_by").
		Updates(models.Pin{Status: pin.Status, AssignedTo: pin.AssignedTo, UsedBy: pin.UsedBy})
	return nil
}
