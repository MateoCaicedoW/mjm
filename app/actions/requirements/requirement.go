package requirements

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/nulls"
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

	requirements := models.Requirements

	// Paginate results. Params "page" and "per_page" control pagination.
	q := pagination(tx, c)

	// Retrieve all Requirements from the DB
	if err := q.Eager().Where("approved_by is null and declined_by is null").All(&requirements); err != nil {

		return err
	}
	path := c.Value("current_route").(buffalo.RouteInfo)

	counters(tx, c)
	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	c.Set("path", path.PathName)
	c.Set("requirements", requirements)

	return c.Render(http.StatusOK, r.HTML("requirement/index.plush.html"))
}

// Show gets the data for one Requirement. This function is mapped to
// the path GET /requirements/{requirement_id}
func Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Requirement
	requirement := &models.Requirement{}

	// To find the Requirement the parameter requirement_id is used.
	if err := tx.Eager().Find(requirement, c.Param("requirement_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("requirement", requirement)

	return c.Render(http.StatusOK, r.HTML("/requirement/show.plush.html"))
}

// New renders the form for creating a new Requirement.
// This function is mapped to the path GET /requirements/new
func New(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	requirement := &models.Requirement{}

	if !ok {
		return fmt.Errorf("no transaction found")
	}

	//set all dropdown
	setDropdowns(tx, c)

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
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	requirement.CreatedByUserID = uuid.FromStringOrNil("175afda1-82ef-4950-b8db-6dab15740d63")
	requirement.RequestingDepartmentID = uuid.FromStringOrNil("668eca48-bc11-49ff-81ea-2665d3130b42")
	requirement.RequirementSubTypeID = uuid.FromStringOrNil("ac5cbfda-8cc8-4d10-8a03-d55ed2647d2d")
	// Validate the data from the html form
	verrs, err := tx.Eager().ValidateAndCreate(requirement)
	if err != nil {

		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		c.Set("requirement", requirement)
		setDropdowns(tx, c)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/requirement/new.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "requirement was successfully created")

	// and redirect to the show page
	return c.Redirect(http.StatusSeeOther, "/requirements")
}

// // Edit renders a edit form for a Requirement. This function is
// // mapped to the path GET /requirements/{requirement_id}/edit
func Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	setDropdowns(tx, c)
	// Allocate an empty Requirement
	requirement := &models.Requirement{}

	if err := tx.Find(requirement, c.Param("requirement_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("requirement", requirement)

	return c.Render(http.StatusOK, r.HTML("/requirement/edit.plush.html"))
}

// // Update changes a Requirement in the DB. This function is mapped to
// // the path PUT /requirements/{requirement_id}
func Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Requirement
	requirement := &models.Requirement{}

	if err := tx.Find(requirement, c.Param("requirement_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Requirement to the html form elements
	if err := c.Bind(requirement); err != nil {
		return err
	}

	verrs, err := tx.Eager().ValidateAndUpdate(requirement)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		setDropdowns(tx, c)
		// Render again the edit.html template that the user can
		// correct the input.
		c.Set("requirement", requirement)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("/requirement/edit.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "requirement was successfully updated")

	// and redirect to the show page
	return c.Redirect(http.StatusSeeOther, "/requirements")
}

// Destroy deletes a Requirement from the DB. This function is mapped
// to the path DELETE /requirements/{requirement_id}
func Delete(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Requirement
	requirement := &models.Requirement{}

	// To find the Requirement the parameter requirement_id is used.
	if err := tx.Find(requirement, c.Param("requirement_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Eager().Destroy(requirement); err != nil {
		return err
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "requirement was successfully deleted")

	// Redirect to the index page
	return c.Redirect(http.StatusSeeOther, "/requirements")
}

func Approved(c buffalo.Context) error {
	// Get the DB connection from the context
	showPendingApprovedDenied("is not null", "approved", c)

	return c.Render(http.StatusOK, r.HTML("/requirement/index.plush.html"))
}

func Pending(c buffalo.Context) error {
	// Get the DB connection from the context
	showPendingApprovedDenied("is null", "pending", c)
	return c.Render(http.StatusOK, r.HTML("/requirement/index.plush.html"))
}

func Denied(c buffalo.Context) error {
	// Get the DB connection from the context
	showPendingApprovedDenied("is not null", "denied", c)
	return c.Render(http.StatusOK, r.HTML("/requirement/index.plush.html"))
}

func Approve(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	// Allocate an empty Requirement
	requirement := &models.Requirement{}

	if err := tx.Find(requirement, c.Param("requirement_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}
	requirement.ApprovedByUserID = nulls.NewUUID(uuid.Must(uuid.FromString("175afda1-82ef-4950-b8db-6dab15740d63")))
	requirement.ApprovedAt = nulls.Time{Time: time.Now(), Valid: true}
	if err := tx.Update(requirement); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/requirements")
}

func Deny(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	// Allocate an empty Requirement
	requirement := &models.Requirement{}

	if err := tx.Find(requirement, c.Param("requirement_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}
	requirement.DeclinedByUserID = nulls.NewUUID(uuid.Must(uuid.FromString("175afda1-82ef-4950-b8db-6dab15740d63")))
	requirement.DeclinedAt = nulls.Time{Time: time.Now(), Valid: true}
	if err := tx.Update(requirement); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/requirements")
}
func setDropdownUsers(users models.Users, c buffalo.Context) {
	userMap := users.Map()

	userMap["Select an User"] = uuid.Nil
	c.Set("users", userMap)
}

func setDropdownDepartment(departments models.Departments, c buffalo.Context) {
	departmentMap := departments.Map()
	departmentMap["Select an Area"] = uuid.Nil
	c.Set("departments", departmentMap)
}

func setDropdownRequirementType(types models.RequirementTypes, c buffalo.Context) {
	typeMap := types.Map()

	typeMap["Select a Type"] = uuid.Nil
	c.Set("requirementTypes", typeMap)

}

func setDropdownRequirementSubType(subtypes models.RequirementSubTypes, c buffalo.Context) {
	subtypeMap := subtypes.Map()
	subtypeMap["Select a Subtype"] = uuid.Nil
	c.Set("requirementSubTypes", subtypeMap)

}

func setDropdowns(tx *pop.Connection, c buffalo.Context) error {
	users := []models.User{}
	if err := tx.All(&users); err != nil {
		return err
	}
	setDropdownUsers(users, c)
	//create map of departments
	departments := []models.Department{}
	if err := tx.All(&departments); err != nil {
		return err
	}
	setDropdownDepartment(departments, c)

	//create map of requirement types
	requirementTypes := []models.RequirementType{}
	if err := tx.All(&requirementTypes); err != nil {
		return err
	}
	setDropdownRequirementType(requirementTypes, c)

	//create map of requirement sub types
	requirementSubType := []models.RequirementSubType{}
	if err := tx.All(&requirementSubType); err != nil {
		return err
	}
	setDropdownRequirementSubType(requirementSubType, c)
	return nil
}

func pagination(tx *pop.Connection, c buffalo.Context) *pop.Query {
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())
	q = q.Order("created_at desc")
	return q
}

func showPendingApprovedDenied(where string, option string, c buffalo.Context) {
	tx, _ := c.Value("tx").(*pop.Connection)

	// Allocate an empty Requirement
	requirements := models.Requirements
	q := pagination(tx, c)
	if option == "approved" {
		condition := "approved_by " + where
		q.Eager().Where(condition).All(&requirements)

	}

	if option == "pending" {
		condition := "approved_by " + where + " and declined_by is null"
		q.Eager().Where(condition).All(&requirements)

	}
	if option == "denied" {
		condition := "declined_by " + where
		q.Eager().Where(condition).All(&requirements)

	}
	counters(tx, c)
	c.Set("requirements", requirements)
	c.Set("pagination", q.Paginator)
	path := c.Value("current_route").(buffalo.RouteInfo)
	c.Set("path", path.PathName)
}

func counters(tx *pop.Connection, c buffalo.Context) {
	// Allocate an empty Requirement
	requirement := &models.Requirements
	pending, _ := tx.Where("approved_by is null and declined_by is null").Count(requirement)
	approved, _ := tx.Where("approved_by is not null").Count(requirement)
	denied, _ := tx.Where("declined_by is not null").Count(requirement)
	c.Set("pending", pending)
	c.Set("approved", approved)
	c.Set("denied", denied)
}
