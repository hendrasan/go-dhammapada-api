package database

import (
	"fmt"
	"log"

	"github.com/hendrasan/go-dhammapada-api/config"
	"github.com/hendrasan/go-dhammapada-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	log.Printf("Connecting to database with DSN: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Verify connected database
	var dbName string
	if err := db.Raw("SELECT current_database()").Scan(&dbName).Error; err != nil {
		return nil, fmt.Errorf("failed to verify database: %v", err)
	}
	log.Printf("Connected to database: %s", dbName)

	return db, nil
}

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Chapter{}, &models.Verse{})
	if err != nil {
		return fmt.Errorf("failed to auto migrate database: %v", err)
	}
	return nil
}
