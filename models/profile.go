package models

import (
	"time"

	"github.com/google/uuid"
)

// Individual model
type Individual struct {
	ID                      uuid.UUID  `json:"id" gorm:"id;primaryKey;not null;type:uuid"`
	HereditaryTitle         string     `json:"hereditary_title"`
	EducationalTitle        string     `json:"educational_title"`
	ProfesionalTitle        string     `json:"professional_title"`
	FirstName               string     `json:"first_name"`
	MiddleName              string     `json:"middle_name"`
	LastName                string     `json:"last_name"`
	PopularName             string     `json:"popular_name"`
	NickName                string     `json:"nick_name"`
	Gender                  string     `json:"gender"`
	DOBActual               string     `json:"dob_actual"`
	DOBVirtual              string     `json:"dob_virtual"`
	AddressMap              []*Address `json:"address_map"  gorm:"foreignKey:individual_id"`
	FamilySpouseID          *uuid.UUID `json:"family_spouse_id"`
	FamilySpouse            *Family    `json:"family_spouse" gorm:"foreignKey:family_spouse_id"`
	FamilyBiologicalChildID *uuid.UUID `json:"family_biological_child_id"`
	FamilyBiologicalChild   *Family    `json:"family_biological_child" gorm:"foreignKey:family_biological_child_id"`
	FamilyAdoptedChildID    *uuid.UUID `json:"family_adopted_child_id"`
	FamilyAdoptedChild      *Family    `json:"family_adopted_child" gorm:"foreignKey:family_adopted_child_id"`
	CreatedAt               time.Time  `json:"created_at"`
	UpdatedAt               time.Time  `json:"updated_at"`
}


