package requirement_type

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gofrs/uuid"

	"github.com/gobuffalo/pop/v6"

	"mjm/app/models"
	"mjm/app/render"
)

var (
	r = render.Engine
)

func List(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	requirementTypes := &models.RequirementTypes{}

	q := tx.PaginateFromParams(c.Params())

	if err := q.All(requirementTypes); err != nil {
		return err
	}

	c.Set("pagination", q.Paginator)
	c.Set("requirementTypes", requirementTypes)

	return c.Render(http.StatusOK, r.HTML("/requirement_type/index.plush.html"))
}

func Show(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	requirementType := &models.RequirementType{}

	if err := tx.Find(requirementType, c.Param("requirement_type_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("requirementType", requirementType)

	return c.Render(http.StatusOK, r.HTML("/requirement_type/show.plush.html"))
}

func New(c buffalo.Context) error {
	c.Set("requirementType", &models.RequirementType{})

	return c.Render(http.StatusOK, r.HTML("/requirement_type/new.plush.html"))
}

func Create(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	requirementType := &models.RequirementType{}

	if err := c.Bind(requirementType); err != nil {
		return err
	}

	requirementType.CreatedByUserID = uuid.Must(uuid.FromString("175afda1-82ef-4950-b8db-6dab15740d63"))

	verrs, err := tx.ValidateAndCreate(requirementType)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("requirementType", requirementType)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/requirement_type/new.plush.html"))
	}

	c.Flash().Add("success", "requirementType.created.success")

	return c.Redirect(http.StatusSeeOther, "requirementTypesPath()")
}

func Edit(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	requirementType := &models.RequirementType{}

	if err := tx.Find(requirementType, c.Param("requirement_type_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("requirementType", requirementType)

	return c.Render(http.StatusOK, r.HTML("/requirement_type/edit.plush.html"))
}

func Update(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	requirementType := &models.RequirementType{}

	if err := tx.Find(requirementType, c.Param("requirement_type_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := c.Bind(requirementType); err != nil {
		return err
	}

	 requirementType.CreatedByUserID = uuid.Must(uuid.FromString("175afda1-82ef-4950-b8db-6dab15740d63"))

	verrs, err := tx.ValidateAndUpdate(requirementType)
	if err != nil {
		return err
	}

	if verrs.HasAny() {

		c.Set("errors", verrs)
		c.Set("requirementType", requirementType)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/requirement_type/edit.plush.html"))
	}

	c.Flash().Add("success", "requirementType.updated.success")

	return c.Redirect(http.StatusSeeOther, "requirementTypesPath()")
}

func Delete(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	requirementType := &models.RequirementType{}

	if err := tx.Find(requirementType, c.Param("requirement_type_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(requirementType); err != nil {
		return err
	}

	c.Flash().Add("success", "requirementType.destroyed.success")

	return c.Redirect(http.StatusSeeOther, "requirementTypesPath()")
}
