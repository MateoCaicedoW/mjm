package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/pop/v6"

	"mjm/app/models"
)

// List gets all RequirementTypes. This function is mapped to the path
// GET /requirement_types
func List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	requirementTypes := &models.RequirementTypes{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all RequirementTypes from the DB
	if err := q.All(requirementTypes); err != nil {
		return err
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	c.Set("requirementTypes", requirementTypes)

	return c.Render(http.StatusOK, r.HTML("/requirement_type/index.plush.html"))
}

// Show gets the data for one RequirementType. This function is mapped to
// the path GET /requirement_types/{requirement_type_id}
func Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty RequirementType
	requirementType := &models.RequirementType{}

	// To find the RequirementType the parameter requirement_type_id is used.
	if err := tx.Find(requirementType, c.Param("requirement_type_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("requirementType", requirementType)

	return c.Render(http.StatusOK, r.HTML("/requirement_types/show.plush.html"))
}

// New renders the form for creating a new RequirementType.
// This function is mapped to the path GET /requirement_types/new
func New(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	departments := models.Departments{}
	if err := tx.All(&departments); err != nil {
		return err
	}
	c.Set("departments", departments.Map())
	c.Set("requirementType", &models.RequirementType{})

	return c.Render(http.StatusOK, r.HTML("/requirement_type/new.plush.html"))
}

// Create adds a RequirementType to the DB. This function is mapped to the
// path POST /requirement_types
func Create(c buffalo.Context) error {

	requirementType := &models.RequirementType{}

	if err := c.Bind(requirementType); err != nil {
		return err
	}

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

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

	return c.Redirect(http.StatusSeeOther, "requirementTypePath()")
}

// Edit renders a edit form for a RequirementType. This function is
// mapped to the path GET /requirement_types/{requirement_type_id}/edit
func Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty RequirementType
	requirementType := &models.RequirementType{}

	if err := tx.Find(requirementType, c.Param("requirement_type_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("requirementType", requirementType)

	return c.Render(http.StatusOK, r.HTML("/requirement_types/edit.plush.html"))
}

// Update changes a RequirementType in the DB. This function is mapped to
// the path PUT /requirement_types/{requirement_type_id}
func Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty RequirementType
	requirementType := &models.RequirementType{}

	if err := tx.Find(requirementType, c.Param("requirement_type_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind RequirementType to the html form elements
	if err := c.Bind(requirementType); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(requirementType)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		c.Set("requirementType", requirementType)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/requirement_types/edit.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "requirementType.updated.success")

	// and redirect to the show page
	return c.Redirect(http.StatusSeeOther, "requirementTypePath()", render.Data{"requirement_type_id": requirementType.ID})
}

// Destroy deletes a RequirementType from the DB. This function is mapped
// to the path DELETE /requirement_types/{requirement_type_id}
func Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty RequirementType
	requirementType := &models.RequirementType{}

	// To find the RequirementType the parameter requirement_type_id is used.
	if err := tx.Find(requirementType, c.Param("requirement_type_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(requirementType); err != nil {
		return err
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "requirementType.destroyed.success")

	// Redirect to the index page
	return c.Redirect(http.StatusSeeOther, "requirementTypesPath()")
}
