package routes

import (
	"bytes"
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/pocketbase/pocketbase/core"
)

func sendPage(e *core.RequestEvent, page templ.Component) error {
	buf := new(bytes.Buffer)
	if err := page.Render(context.Background(), buf); err != nil {
		return err
	}

	return e.HTML(http.StatusOK, buf.String())
}
