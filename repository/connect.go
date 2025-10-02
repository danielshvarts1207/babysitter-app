package repository

import (
	"babysitter-app/models/entities"
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(connectionString string) (*gorm.DB, context.Context) {
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}

	ctx := context.Background()

	db.AutoMigrate(&entities.Babysitter{})

	return db, ctx
}
