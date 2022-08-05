package requirement

import (
	"mjm/app/models"
	"mjm/app/render"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

var (
	r = render.Engine
)

func New(c buffalo.Context) error {
	requirementsTypes := models.RequirementTypes{}
	requirementsSubTypes := models.RequirementSubTypes{}

	subtypesOptions := map[string]uuid.UUID{}
	subtypesOptions["Select  a Subtype"] = uuid.Nil

	serviceArea := map[string]uuid.UUID{}
	serviceArea["Select  an area"] = uuid.Nil

	c.Set("subtypesOptions", requirementsSubTypes.Map())
	c.Set("serviceArea", serviceArea)
	c.Set("typeOptions", requirementsTypes.Map())

	return c.Render(http.StatusOK, r.HTML("/requirement/new.plush.html"))
}

func Create(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	requirement := models.Requirement{}
	if err := c.Bind(&requirement); err != nil {
		return err
	}

	errCreate := tx.Create(&requirement)
	if errCreate != nil {
		return errCreate
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
