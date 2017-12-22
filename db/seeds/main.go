package main

import (
	//	"encoding/json"
	"fmt"
	//	"io"
	//	"log"
	//	"math/rand"
	//	"net/http"
	//	"net/url"
	//	"os"
	//	"path/filepath"
	//	"regexp"
	//	"strconv"
	//	"strings"
	"time"

	//	"github.com/jinzhu/now"
	"github.com/linkonoid/qoradmin/app/models"
	//	"github.com/linkonoid/qoradmin/config/admin"
	"github.com/linkonoid/qoradmin/config/auth"
	//	"github.com/linkonoid/qoradmin/db"
	_ "github.com/linkonoid/qoradmin/db/migrations"
	"github.com/qor/auth/auth_identity"
	"github.com/qor/auth/providers/password"
	i18n_database "github.com/qor/i18n/backends/database"
	"github.com/qor/notification"
	//	"github.com/qor/notification/channels/database"
	"github.com/qor/qor"
	//	"github.com/qor/slug"
	//	"github.com/qor/sorting"
)

/* How to run this script
   $ go run db/seeds/main.go db/seeds/seeds.go
*/

var (
	AdminUser    *models.User
	Notification = notification.New(&notification.Config{})
	Tables       = []interface{}{
		&auth_identity.AuthIdentity{},
		&models.UserGroup{},
		&models.User{},
		&models.Category{},
		&models.Product{},
		&i18n_database.Translation{},
		&notification.QorNotification{},
	}
)

func createAdminUsers() {
	AdminUser = &models.User{}
	AdminUser.Email = "test@test.com"
	AdminUser.Confirmed = true
	AdminUser.Name = "QOR Admin"
	AdminUser.Role = "Admin"
	DraftDB.Create(AdminUser)

	provider := auth.Auth.GetProvider("password").(*password.Provider)
	hashedPassword, _ := provider.Encryptor.Digest("testing")
	now := time.Now()

	authIdentity := &auth_identity.AuthIdentity{}
	authIdentity.Provider = "password"
	authIdentity.UID = AdminUser.Email
	authIdentity.EncryptedPassword = hashedPassword
	authIdentity.UserID = fmt.Sprint(AdminUser.ID)
	authIdentity.ConfirmedAt = &now

	DraftDB.Create(authIdentity)

	// Send welcome notification
	Notification.Send(&notification.Message{
		From:        AdminUser,
		To:          AdminUser,
		Title:       "Welcome To QOR Admin",
		Body:        "Welcome To QOR Admin",
		MessageType: "info",
	}, &qor.Context{DB: DraftDB})
}

func main() {
	createRecords()
}

func createRecords() {
	createAdminUsers()
	fmt.Println("--> Created admin users.")
}
