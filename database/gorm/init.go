package gorm

import (
	"fmt"
	"synergy/web-service-gin/common/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	config, _ := config.LoadConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
	// open
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println("# gorm database open.")
	return db, err
}

type Customer struct {
	ID       uint      `gorm:"column:customer_id;primaryKey;"`
	Name     string    `gorm:"column:customer_name;"`
	Contacts []Contact `gorm:"foreignkey:CustomerID"`
}

type Contact struct {
	ID         uint `gorm:"column:contact_id;primaryKey;"`
	CustomerID uint `gorm:"column:customer_id;"`
}

func GetAllData() []Customer {
	db, _ := InitDatabase()
	// Read
	var customer []Customer
	db.Table("public.customers").Find(&customer)
	fmt.Printf("# gorm query %d", len(customer))
	return customer
}
