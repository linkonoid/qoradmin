package admin

import (
	"github.com/linkonoid/qoradmin/app/models"
	"github.com/linkonoid/qoradmin/config/i18n"
	"github.com/linkonoid/qoradmin/db"
	_ "github.com/linkonoid/qoradmin/db/migrations"
	"github.com/linkonoid/qoradmin/config/auth"
	"github.com/qor/admin"
)

var Admin *admin.Admin

func init() {

	Admin = admin.New(&admin.AdminConfig{DB: db.DB, Auth: auth.AdminAuth{}, SiteName: "qoradmin Example"})

	Admin.AddResource(&models.UserGroup{}, &admin.Config{Menu: []string{"User Management"}, Name: "UserGroup"})
	Admin.AddResource(&models.User{}, &admin.Config{Menu: []string{"User Management"}, Name: "User"})
	Admin.AddResource(&models.Category{}, &admin.Config{Menu: []string{"Product Management"}})
	Admin.AddResource(&models.Product{}, &admin.Config{Menu: []string{"Product Management"}})

	Admin.AddResource(i18n.I18n)

}

