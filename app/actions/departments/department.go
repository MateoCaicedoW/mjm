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

func List(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	departments := []models.Department{}
	err := tx.All(&departments)
	if err != nil {
		return err
	}

	c.Set("departments", departments)

	return c.Render(http.StatusOK, r.HTML("/departments/index.plush.html"))
}

func New(c buffalo.Context) error {
	c.Set("department", &models.Department{})

	return c.Render(http.StatusOK, r.HTML("/departments/new.plush.html"))
}

func Create(c buffalo.Context) error {
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

func Show(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	department := models.Department{}
	depatmentID := c.Param("department_id")

	err := tx.Find(&department, depatmentID)
	if err != nil {
		return err
	}

	c.Set("department", department)

	return c.Render(http.StatusOK, r.HTML("/departments/show.plush.html"))
}

func Edit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	department := models.Department{}
	departmentID := c.Param("department_id")

	err := tx.Find(&department, departmentID)
	if err != nil {
		return err
	}

	c.Set("department", department)

	return c.Render(http.StatusOK, r.HTML("/departments/edit.plush.html"))
}

func Update(c buffalo.Context) error {
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

func Delete(c buffalo.Context) error {
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
