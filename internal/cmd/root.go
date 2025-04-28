package cmd

import (
	"goth/internal/config"
	"goth/internal/routes"
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "goth/internal/migrations" // register all migrations
)

func Root() {
	app := pocketbase.New()
	cfg := config.GetServerConfigs()
	app.OnServe().BindFunc(routes.SetupRoutes)
	app.Settings().Meta.HideControls = cfg.IsProd

	if !cfg.IsProd {
		migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{})
	}

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
