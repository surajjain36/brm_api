package postgres

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// Connection is a pointer to gorm.DB
	Connection *gorm.DB
)

func NewConnect() (err error) {
	err = connectMoolDB()
	if err != nil {
		return err
	}

	return nil
}

// NewMoolDBConnect with database
func connectMoolDB() (err error) {
	var (
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		host     = os.Getenv("DB_HOST")
		db       = os.Getenv("DB_NAME")
		port     = os.Getenv("DB_PORT")
	)

	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%s sslmode=disable", user, password, host, db, port)
	Connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := Connection.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Second)

	fmt.Println("Connected with Database")
	return nil
}
