package migrations

import (
	"brm_api/services/db/postgres"
	"brm_api/services/userservice"
	"brm_api/utils/common"
	"time"

	"brm_api/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type family struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;not null;uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type individual struct {
	ID               uuid.UUID `json:"id" gorm:"id;primaryKey;not null;type:uuid"`
	Initials         string    `json:"initials"`
	HereditaryTitle  string    `json:"hereditary_title"`
	EdicationalTitle string    `json:"educational_title"`
	ProfesionalTitle string    `json:"professional_title"`
	FirstName        string    `json:"first_name"`
	MiddleName       string    `json:"middle_name"`
	LastName         string    `json:"last_name"`
	PopularName      string    `json:"popular_name"`
	NickName         string    `json:"nick_name"`
	Gender           string    `json:"gender"`
	DOBActual        string    `json:"dob_actual"`
	DOBVirtual       string    `json:"dob_virtual"`
	Mobile           string    `json:"mobile" gorm:"mobile;index;not null" validate:"required"`
	CountryCode      string    `json:"country_code" gorm:"index;not null" validate:"required"`

	FamilySpouseID uuid.UUID `json:"family_spouse_id"`
	FamilySpouse   *family   `gorm:"foreignKey:family_spouse_id"`

	FamilyBiologicalChildID uuid.UUID `json:"family_biological_child_id"`
	FamilyBiologicalChild   *family   `gorm:"foreignKey:family_biological_child_id"`

	FamilyAdoptedChildID uuid.UUID `json:"family_adopted_child_id"`
	FamilyAdoptedChild   *family   `gorm:"foreignKey:family_adopted_child_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// func creatFamilyBaseStructure() {
// 	db := postgres.Connection
// 	db.AutoMigrate(&family{}, &individual{})
// }

func createDependencyTables() {
	db := postgres.Connection
	roles := []models.Role{
		{ID: 1, Name: "Super Admin", Key: "super_admin"},
		{ID: 2, Name: "Admin", Key: "admin"},
		{ID: 3, Name: "Customer", Key: "customer"},
	}
	packages := []models.Package{
		{Name: "Achievers club", Key: models.AchieverKey, Amount: 10900},
		{Name: "3990 club", Key: models.ClassicKey, Amount: 3990},
	}
	db.AutoMigrate(&roles, &packages)

	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Role{})
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},                    // key colume
		DoUpdates: clause.AssignmentColumns([]string{"name", "key"}), // column needed to be updated
	}).Create(&roles)

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},                              // key colume
		DoUpdates: clause.AssignmentColumns([]string{"name", "key", "amount"}), // column needed to be updated
	}).Create(&packages)
}

func createSuperAdmin() {
	user := models.User{
		MobileNumber: "9964582028",
		Password:     common.HashAndSalt("9964582028"),
		Role:         models.SuperAdmin,
	}
	userservice.CreateUser(user)
}

func createUserRelationsTables() {
	db := postgres.Connection
	db.Debug().Migrator().CreateTable(&models.Pin{}, &models.PinTransaction{})
	db.Debug().Migrator().CreateTable(&models.User{}, &models.UserPackage{})
}
