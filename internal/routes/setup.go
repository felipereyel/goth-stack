package routes

import (
	"goth/internal/embeded"

	"github.com/pocketbase/pocketbase/core"

	"github.com/pocketbase/pocketbase/apis"
)

var assetsHandler = apis.Static(embeded.Assets, true)

func healthzHandler(e *core.RequestEvent) error {
	return e.String(200, "ok")
}

func SetupRoutes(se *core.ServeEvent) error {
	se.Router.GET("/", taskList)
	se.Router.GET("/new", taskNew)
	se.Router.GET("/edit/{id}", taskEdit)
	se.Router.POST("/edit/{id}", taskSave)

	se.Router.GET("/healthz", healthzHandler)
	se.Router.GET("/statics/{path...}", assetsHandler)

	return se.Next()
}
