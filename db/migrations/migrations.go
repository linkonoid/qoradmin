package migrations

import (
	"github.com/linkonoid/qoradmin/app/models"
	"github.com/linkonoid/qoradmin/db"
)

func init() {
	db.DB.AutoMigrate(&models.UserGroup{}, &models.User{})
	db.DB.AutoMigrate(&models.Category{}, &models.Product{})
}

func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.DB.AutoMigrate(value)
	}
}
