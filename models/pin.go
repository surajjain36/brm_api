package models

import "time"

var (
	UnUsedStatus = "un_used"
	UsedStatus   = "used"
)

type Pin struct {
	ID          uint   `gorm:"id;primaryKey;not null;autoIncrement" json:"id"`
	Pin         string `json:"pin"`
	Status      string `json:"status"`
	GeneratedBy uint   `json:"generated_by"`
	PackageID   uint   `json:"package_id"`
	// ShareTo     uint      `json:"share_to"`
	// SharedBy    uint      `json:"shared_by"` //if the pins are transfered.
	UsedBy     uint      `json:"used_by"`     //Person who used this pin to add some other user
	AssignedTo uint      `json:"assigned_to"` //Person who uses the pin to get added.
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type PinTransaction struct {
	ID uint `gorm:"id;primaryKey;not null;autoIncrement" json:"id"`
	//Pin         string    `json:"pin"`
	//Status      string    `json:"status"`
	//GeneratedBy uint      `json:"generated_by"`
	PinID uint `json:"pin_id" gorm:"foreignkey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	//Pin           //  Pin  `json:"pin"`
	SharedTo uint `json:"shared_to"`
	SharedBy uint `json:"shared_by"` //if the pins are transfered.
	//UsedBy   uint `json:"used_by"`
	//AssignedTo uint      `json:"assigned_to"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
