package main

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sportshot/pkg/utils/db"
	"sportshot/pkg/utils/global"
	"sportshot/pkg/utils/models/user"
)

func main() {

	global.MySQLClient = db.GetMySQLClient()

	err := global.MySQLClient.AutoMigrate(&user.User{})
	if err != nil {
		panic("failed to migrate database")
	}

	global.MySQLClient.Create(&user.User{
		Model:        gorm.Model{},
		ID:           uuid.UUID{},
		Username:     "123",
		Password:     "123",
		RefreshToken: "123",
		Active:       false,
	})

}
