package models

import (
	"time"
)

// Context Keys
const (
	UserKey = "user"
)

var (
	//Roles
	SuperAdmin = "super_admin"
	Admin      = "admin"
	// Manager    = "manager"
	// Leader     = "leader"
	Customer = "customer"
	//Packages
	AchieverKey = "achievers"
	ClassicKey  = "classic"
)

//User model
type User struct {
	ID           uint      `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	EmailID      string    `json:"email_id"`
	MobileNumber string    `json:"mobile_number" gorm:"unique;index;not null" validate:"required"`
	OTP          string    `json:"-" gorm:"otp;size:6"`
	Password     string    `json:"-"`
	OTPCreatedAt time.Time `json:"-"`
	DeviceID     string    `json:"-"`
	AccessToken  string    `json:"-"`
	Profile                //Profile related fields
	Hierarchy    string    `json:"-"`
	Role         string    `json:"role"`
	//PinID        uint      `json:"pin_id" gorm:"default:null"`
	//Pin          Pin       `json:"pin"`
	//Pins      []Pin         `json:"pins" gorm:"foreignkey:SharedTo;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Packages  []UserPackage `json:"packages" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PartnerID uint          `json:"partner_id" gorm:"default:null"`
	Partner   []User        `json:"partner" gorm:"foreignkey:PartnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RefererID uint          `json:"referer_id"`
	CreatedBy uint          `json:"created_by"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	DeletedAt *time.Time    `json:"deleted_at"`
}

type UserPackage struct {
	ID        uint   `json:"id" gorm:"id;primaryKey;not null;autoIncrement"`
	UserID    uint   `json:"user_id"`
	Package   string `json:"package"`
	PackageID uint   `json:"package_id" gorm:"foreignkey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PinID     uint   `json:"pin_id" gorm:"foreignkey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Profile model
type Profile struct {
	Name   string `json:"name" validate:"required"`
	DOB    string `json:"dob" validate:"required"`
	Gender string `json:"gender"`
	Address
	BankAccount
	PanNumber           string `json:"pan_number"`
	NomineeName         string `json:"nominee_name"`
	NomineeContactNo    string `json:"nominee_contact_no"`
	NomineeRelationship string `json:"nominee_relationship"`
	ProfilePictureLink  string `json:"profile_picture"`
}

// Address model
type Address struct {
	Address  string `json:"address" validate:"required"`
	Address1 string `json:"address1"`
	Village  string `json:"village"`
	District string `json:"district" validate:"required"`
	State    string `json:"state" validate:"required"`
	Pincode  string `json:"pincode" validate:"required"`
	Country  string `json:"country"`
}

// BankAccount model
type BankAccount struct {
	AccountNumber string `json:"account_number" validate:"required"`
	IFSCCode      string `json:"ifsc_code" validate:"required"`
	BankName      string `json:"bank_name"`
	BranchName    string `json:"branch_name"`
}

//Role model
type Role struct {
	ID   uint   `gorm:"type:primaryKey;not null;autoIncrement" json:"id"`
	Name string `json:"name"`
	Key  string `json:"key" gorm:"unique;not null"`
}

type Package struct {
	ID          uint   `gorm:"type:primaryKey;not null;autoIncrement" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Key         string `json:"key" gorm:"unique;not null"`
}
