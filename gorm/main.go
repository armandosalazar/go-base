package gorm

import (
	"go-base/gorm/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Run() {
	db, err := gorm.Open(
		sqlite.Open("test.db"),
		&gorm.Config{},
	)

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}, &models.User{})

	db.Create(&Product{Code: "D42", Price: 100})
	db.Create(&models.User{Name: "Jinzhu", Email: "jinzhu@email.com"})

}
