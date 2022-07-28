package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"

	"mjm/app/models"
)

// RequirementsResource is the resource for the Requirement model
type RequirementsResource struct {
	buffalo.Resource
}

// List gets all Requirements. This function is mapped to the path
// GET /requirements
func (v RequirementsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	requirements := &models.Requirements

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Requirements from the DB
	if err := q.All(requirements); err != nil {
		return err
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	c.Set("requirements", requirements)

	return c.Render(http.StatusOK, r.HTML("/requirements/index.plush.html"))
}

// Show gets the data for one Requirement. This function is mapped to
// the path GET /requirements/{requirement_id}
// func (v RequirementsResource) Show(c buffalo.Context) error {
// 	// Get the DB connection from the context
// 	tx, ok := c.Value("tx").(*pop.Connection)
// 	if !ok {
// 		return fmt.Errorf("no transaction found")
// 	}

// 	// Allocate an empty Requirement
// 	requirement := &models.Requirement{}

// 	// To find the Requirement the parameter requirement_id is used.
// 	if err := tx.Find(requirement, c.Param("requirement_id")); err != nil {
// 		return c.Error(http.StatusNotFound, err)
// 	}

// 	c.Set("requirement", requirement)

// 	return c.Render(http.StatusOK, r.HTML("/requirements/show.plush.html"))
// }

// New renders the form for creating a new Requirement.
// This function is mapped to the path GET /requirements/new
// func (v RequirementsResource) New(c buffalo.Context) error {
// 	c.Set("requirement", &models.Requirement{})

// 	return c.Render(http.StatusOK, r.HTML("/requirements/new.plush.html"))
// }

// // Create adds a Requirement to the DB. This function is mapped to the
// // path POST /requirements
// func (v RequirementsResource) Create(c buffalo.Context) error {
// 	// Allocate an empty Requirement
// 	requirement := &models.Requirement{}

// 	// Bind requirement to the html form elements
// 	if err := c.Bind(requirement); err != nil {
// 		return err
// 	}

// 	// Get the DB connection from the context
// 	tx, ok := c.Value("tx").(*pop.Connection)
// 	if !ok {
// 		return fmt.Errorf("no transaction found")
// 	}

// 	// Validate the data from the html form
// 	verrs, err := tx.ValidateAndCreate(requirement)
// 	if err != nil {
// 		return err
// 	}

// 	if verrs.HasAny() {
// 		// Make the errors available inside the html template
// 		c.Set("errors", verrs)

// 		// Render again the new.html template that the user can
// 		// correct the input.
// 		c.Set("requirement", requirement)

// 		return c.Render(http.StatusUnprocessableEntity, r.HTML("/requirements/new.plush.html"))
// 	}

// 	// If there are no errors set a success message
// 	c.Flash().Add("success", "requirement.created.success")

// 	// and redirect to the show page
// 	return c.Redirect(http.StatusSeeOther, "requirementPath()", render.Data{"requirement_id": requirement.ID})
// }

// // Edit renders a edit form for a Requirement. This function is
// // mapped to the path GET /requirements/{requirement_id}/edit
// func (v RequirementsResource) Edit(c buffalo.Context) error {
// 	// Get the DB connection from the context
// 	tx, ok := c.Value("tx").(*pop.Connection)
// 	if !ok {
// 		return fmt.Errorf("no transaction found")
// 	}

// 	// Allocate an empty Requirement
// 	requirement := &models.Requirement{}

// 	if err := tx.Find(requirement, c.Param("requirement_id")); err != nil {
// 		return c.Error(http.StatusNotFound, err)
// 	}

// 	c.Set("requirement", requirement)

// 	return c.Render(http.StatusOK, r.HTML("/requirements/edit.plush.html"))
// }

// // Update changes a Requirement in the DB. This function is mapped to
// // the path PUT /requirements/{requirement_id}
// func (v RequirementsResource) Update(c buffalo.Context) error {
// 	// Get the DB connection from the context
// 	tx, ok := c.Value("tx").(*pop.Connection)
// 	if !ok {
// 		return fmt.Errorf("no transaction found")
// 	}

// 	// Allocate an empty Requirement
// 	requirement := &models.Requirement{}

// 	if err := tx.Find(requirement, c.Param("requirement_id")); err != nil {
// 		return c.Error(http.StatusNotFound, err)
// 	}

// 	// Bind Requirement to the html form elements
// 	if err := c.Bind(requirement); err != nil {
// 		return err
// 	}

// 	verrs, err := tx.ValidateAndUpdate(requirement)
// 	if err != nil {
// 		return err
// 	}

// 	if verrs.HasAny() {
// 		// Make the errors available inside the html template
// 		c.Set("errors", verrs)

// 		// Render again the edit.html template that the user can
// 		// correct the input.
// 		c.Set("requirement", requirement)

// 		return c.Render(http.StatusUnprocessableEntity, r.HTML("/requirements/edit.plush.html"))
// 	}

// 	// If there are no errors set a success message
// 	c.Flash().Add("success", "requirement.updated.success")

// 	// and redirect to the show page
// 	return c.Redirect(http.StatusSeeOther, "requirementPath()", render.Data{"requirement_id": requirement.ID})
// }

// // Destroy deletes a Requirement from the DB. This function is mapped
// // to the path DELETE /requirements/{requirement_id}
// func (v RequirementsResource) Destroy(c buffalo.Context) error {
// 	// Get the DB connection from the context
// 	tx, ok := c.Value("tx").(*pop.Connection)
// 	if !ok {
// 		return fmt.Errorf("no transaction found")
// 	}

// 	// Allocate an empty Requirement
// 	requirement := &models.Requirement{}

// 	// To find the Requirement the parameter requirement_id is used.
// 	if err := tx.Find(requirement, c.Param("requirement_id")); err != nil {
// 		return c.Error(http.StatusNotFound, err)
// 	}

// 	if err := tx.Destroy(requirement); err != nil {
// 		return err
// 	}

// 	// If there are no errors set a flash message
// 	c.Flash().Add("success", "requirement.destroyed.success")

// 	// Redirect to the index page
// 	return c.Redirect(http.StatusSeeOther, "requirementsPath()")
// }
