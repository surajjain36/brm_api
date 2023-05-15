package models

import (
	"time"

	"github.com/google/uuid"
)

// Family model
type Family struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;not null;uuid"`

	HusbandID uuid.UUID   `json:"husband_id"`
	Husband   *Individual `json:"husband" gorm:"foreignKey:husband_id"`

	WifeID uuid.UUID   `json:"wife_id"`
	Wife   *Individual `json:"wife"  gorm:"foreignKey:wife_id"`

	BiologicalChildren []*Individual `json:"biological_children" gorm:"foreignKey:family_biological_child_id;references:ID"`

	AdoptedChildren []*Individual `json:"adopted_children" gorm:"foreignKey:family_adopted_child_id;references:ID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
