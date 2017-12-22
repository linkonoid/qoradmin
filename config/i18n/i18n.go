package i18n

import (
	"path/filepath"

	"github.com/linkonoid/qoradmin/config"
	"github.com/linkonoid/qoradmin/db"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	"github.com/qor/i18n/backends/yaml"
	_ "github.com/qor/i18n/inline_edit"
)

var I18n *i18n.I18n

func init() {
	I18n = i18n.New(database.New(db.DB), yaml.New(filepath.Join(config.Root, "config/locales")))
}
