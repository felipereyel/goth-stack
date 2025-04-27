package cmd

import (
	"goth/internal/routes"
	"log"

	"github.com/pocketbase/pocketbase"
)

func Root() {
	app := pocketbase.New()
	app.OnServe().BindFunc(routes.SetupRoutes)

	// cfg := config.GetServerConfigs()
	// if cfg.AutoMigrate {
	// 	migrations.Register()
	// }

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
