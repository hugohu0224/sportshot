package db

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"sportshot/pkg/utils/models"
)

func GetMySQLClient() *gorm.DB {
	// connection
	dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.S().Fatalf("failed to connect mysql: %v", err)
	}
	zap.S().Info("success to connect mysql")

	// sync
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		zap.S().Fatalf("failed to migrate database")
	}
	return db
}
