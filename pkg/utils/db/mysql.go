package db

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GetMySQLClient() *gorm.DB {
	dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.S().Fatalf("failed to connect mysql: %v", err)
	}
	zap.S().Info("success to connect mysql")
	return db
}
