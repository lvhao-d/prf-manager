package db

import (
	"fmt"
	"os"
	"prf-manager/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	fmt.Println(os.Getenv(("DB_USER")))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		fmt.Printf("Failed to connect to database: %v\n", err)
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.Warehouse{}, &entity.Record{}, &entity.User{}, &entity.Agency{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	fmt.Println("Database migration completed successfully")
	return nil
}
