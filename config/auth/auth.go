package auth

import (
	"time"

	"github.com/qor/auth"
	"github.com/qor/auth/auth_identity"
	"github.com/qor/auth/authority"
	//	"github.com/qor/auth/providers/facebook"
	//	"github.com/qor/auth/providers/github"
	//	"github.com/qor/auth/providers/google"
	//	"github.com/qor/auth/providers/twitter"
	"github.com/linkonoid/qoradmin/app/models"
	"github.com/linkonoid/qoradmin/config"
	"github.com/linkonoid/qoradmin/db"
	"github.com/qor/auth/providers/password"
	"github.com/qor/auth_themes/clean"
)

var (
	// Auth initialize Auth for Authentication
	Auth = clean.New(&auth.Config{
		DB: db.DB,
		UserModel:  models.User{},
		Redirector: auth.Redirector{RedirectBack: config.RedirectBack},
	})

	// Authority initialize Authority for Authorization
	Authority = authority.New(&authority.Config{
		Auth: Auth,
	})
)

func init() {
	db.DB.AutoMigrate(&auth_identity.AuthIdentity{})
	//	Auth.RegisterProvider(github.New(&config.Config.Github))
	//	Auth.RegisterProvider(google.New(&config.Config.Google))
	//	Auth.RegisterProvider(facebook.New(&config.Config.Facebook))
	//	Auth.RegisterProvider(twitter.New(&config.Config.Twitter))
	Auth.RegisterProvider(password.New(&password.Config{}))
	Authority.Register("logged_in_half_hour", authority.Rule{TimeoutSinceLastLogin: time.Minute * 30})
}
