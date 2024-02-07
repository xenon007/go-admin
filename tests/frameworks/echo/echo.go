package echo

import (
	// add echo adapter
	_ "github.com/xenon007/go-admin/adapter/echo"
	"github.com/xenon007/go-admin/modules/config"
	"github.com/xenon007/go-admin/modules/language"
	"github.com/xenon007/go-admin/plugins/admin/modules/table"

	// add mysql driver
	_ "github.com/xenon007/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "github.com/xenon007/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/xenon007/go-admin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "github.com/xenon007/go-admin/modules/db/drivers/mssql"
	// add adminlte ui theme
	"github.com/GoAdminGroup/themes/adminlte"

	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/xenon007/go-admin/engine"
	"github.com/xenon007/go-admin/plugins/admin"
	"github.com/xenon007/go-admin/plugins/example"
	"github.com/xenon007/go-admin/template"
	"github.com/xenon007/go-admin/template/chartjs"
	"github.com/xenon007/go-admin/tests/tables"
)

func internalHandler() http.Handler {
	e := echo.New()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)
	template.AddComp(chartjs.NewChart())

	examplePlugin := example.NewExample()

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(e); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return e
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	e := echo.New()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(gens)

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfig(&config.Config{
		Databases: dbs,
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.EN,
		IndexUrl:    "/",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}).
		AddPlugins(adminPlugin).Use(e); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return e
}
