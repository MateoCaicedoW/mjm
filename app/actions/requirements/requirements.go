package requirements

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"

	"mjm/app/models"
	"mjm/app/render"
)

var (
	// r is a buffalo/render Engine that will be used by actions
	// on this package to render render HTML or any other formats.
	r = render.Engine
)

// RequirementsResource is the resource for the Requirement model

// List gets all Requirements. This function is mapped to the path
// GET /requirements
func List(c buffalo.Context) error {
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

	return c.Render(http.StatusOK, r.HTML("requirement/index.plush.html"))
}

// Show gets the data for one Requirement. This function is mapped to
// the path GET /requirements/{requirement_id}
// func  Show(c buffalo.Context) error {
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
func New(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	requirement := &models.Requirement{}
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	users := []models.User{}
	if err := tx.All(&users); err != nil {
		return err
	}
	//create map of users
	userMap := make(map[string]uuid.UUID)

	for i := 0; i < len(users); i++ {

		userMap[users[i].FirstName+" "+users[i].LastName] = users[i].ID
	}
	//create map of departments
	departmentMap := make(map[string]uuid.UUID)
	departments := []models.Department{}
	if err := tx.All(&departments); err != nil {
		return err
	}
	for i := 0; i < len(departments); i++ {
		departmentMap[departments[i].Name] = departments[i].ID
	}

	//create map of requirement types
	requirementTypeMap := make(map[string]uuid.UUID)
	requirementTypes := []models.RequirementType{}
	if err := tx.All(&requirementTypes); err != nil {
		return err
	}
	for i := 0; i < len(requirementTypes); i++ {
		requirementTypeMap[requirementTypes[i].Name] = requirementTypes[i].ID
	}

	//create map of requirement sub types
	requirementSubType := []models.RequirementSubType{}
	requirementSubTypeMap := make(map[string]uuid.UUID)
	if err := tx.All(&requirementSubType); err != nil {
		return err
	}
	for i := 0; i < len(requirementSubType); i++ {
		requirementSubTypeMap[requirementSubType[i].Name] = requirementSubType[i].ID
	}

	c.Set("requirementTypes", requirementTypeMap)
	c.Set("users", userMap)
	c.Set("departments", departmentMap)
	c.Set("requirementSubTypes", requirementSubTypeMap)

	requirement.CreatedAt = time.Now()
	c.Set("requirement", requirement)

	return c.Render(http.StatusOK, r.HTML("/requirement/new.plush.html"))
}

// // Create adds a Requirement to the DB. This function is mapped to the
// // path POST /requirements
func Create(c buffalo.Context) error {
	// Allocate an empty Requirement
	requirement := &models.Requirement{}

	// Bind requirement to the html form elements
	if err := c.Bind(requirement); err != nil {
		return err
	}
	fmt.Println(requirement)
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(requirement, "modified_by")
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		c.Set("requirement", requirement)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/requirements/new.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "requirement.created.success")

	// and redirect to the show page
	return c.Redirect(http.StatusSeeOther, "/requirements")
}

// // Edit renders a edit form for a Requirement. This function is
// // mapped to the path GET /requirements/{requirement_id}/edit
// func  Edit(c buffalo.Context) error {
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
// func  Update(c buffalo.Context) error {
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
// func  Destroy(c buffalo.Context) error {
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
