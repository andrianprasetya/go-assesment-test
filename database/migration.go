package database

import (
	"github.com/andrianprasetya/go-assesment-test/internal/model"
	"gorm.io/gorm"
	"log"
)

// MigrateDatabase runs database migrations
func MigrateDatabase(db *gorm.DB) {

	err := db.AutoMigrate(
		&model.User{},
		&model.Transaction{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migration completed successfully")
}
