package departments

import (
	"mjm/app/models"
	"mjm/app/render"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

var (
	r = render.Engine
)

func IndexDepartment(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	departments := []models.Department{}
	err := tx.All(&departments)
	if err != nil {
		return err
	}

	c.Set("department", departments)

	return c.Render(http.StatusOK, r.HTML("/departments/index.plush.html"))
}

func FormCreateDeparment(c buffalo.Context) error {

	return c.Render(http.StatusOK, r.HTML("/departments/create_departments.plush.html"))
}

func ViewDetails(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	department := models.Department{}

	depatmentID := c.Param("department_id")
	err := tx.Find(&department, depatmentID)
	if err != nil {
		return err
	}

	c.Set("department", department)

	return c.Render(http.StatusOK, r.HTML("/departments/view_details.plush.html"))
}

func CreateDepartment(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	deparment := models.Department{}
	if err := c.Bind(&deparment); err != nil {
		return err
	}

	err := tx.Create(&deparment)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/departments")
}

func Edit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	department := models.Department{}
	departmentID := c.Param("department_id")

	err := tx.Find(&department, departmentID)
	if err != nil {
		return err
	}

	c.Set("deparment", department)
	return c.Render(http.StatusOK, r.HTML("/departments/edit_department.plush.html"))
}

func Editdepartment(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	deparment := models.Department{}
	departmentID := c.Param("department_id")

	err := tx.Find(&deparment, departmentID)
	if err != nil {
		return err
	}

	if err := c.Bind(&deparment); err != nil {
		return err
	}

	err = tx.Update(&deparment)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/departments")
}

func DeleteDepartment(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	department := models.Department{}
	departmentID := c.Param("department_id")

	err := tx.Find(&department, departmentID)
	if err != nil {
		return err
	}

	err = tx.Destroy(&department)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/departments")
}
