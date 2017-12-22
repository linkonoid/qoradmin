package models

import (
	"github.com/jinzhu/gorm"
	//	"github.com/qor/l10n"
	"github.com/qor/sorting"
)

// Create a GORM-backend model
type UserGroup struct {
	gorm.Model
	//l10n.Locale
	sorting.Sorting
	Name string //`l10n:"sync"`
}

type User struct {
	gorm.Model
	//l10n.Locale
	sorting.Sorting
	Group     []UserGroup
	Name      string
	Role      string
	Email     string
	Confirmed bool
}

type Category struct {
	gorm.Model
	//l10n.Locale
	Name        string
	Description string
}

type Product struct {
	gorm.Model
	//l10n.Locale
	Category    Category
	Name        string
	Description string
}

func (user User) AvailableLocales() []string {
	return []string{"en-EN", "ru-RU"}
}

func (user User) DisplayName() string {
	return user.Email
}
