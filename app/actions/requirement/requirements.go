package requirements

import (
	"mjm/app/render"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

var (
	r = render.Engine
)

func Edit(c buffalo.Context) error {

	return c.Render(http.StatusOK, r.HTML("/requirement/edit.plush.html"))
}
