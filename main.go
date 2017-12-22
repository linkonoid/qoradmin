package main

import (
	"fmt"
	"net/http"

	"github.com/linkonoid/qoradmin/config"
	"github.com/linkonoid/qoradmin/config/admin"
	"github.com/linkonoid/qoradmin/config/routes"
	"github.com/qor/middlewares"
)

func main() {
	mux := http.NewServeMux()
	admin.Admin.MountTo("/admin", mux)
	mux.Handle("/", routes.Router())
	fmt.Printf("Listening on: %v\n", config.Config.Port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), middlewares.Apply(mux)); err != nil {
		panic(err)
	}
}
