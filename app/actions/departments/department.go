package departments

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

func List(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	departments := []models.Department{}
	err := tx.All(&departments)
	if err != nil {
		return err
	}

	areaRequirements := []models.AreaRequirementType{}
	if err := tx.Eager().All(&areaRequirements); err != nil {
		return err
	}

	c.Set("requirements", areaRequirements)
	c.Set("departments", departments)

	return c.Render(http.StatusOK, r.HTML("/department/index.plush.html"))
}

func New(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	c.Set("department", &models.Department{})

	requirements := models.RequirementTypes{}
	if err := tx.All(&requirements); err != nil {
		return err
	}

	c.Set("requirements", requirements.Map())
	return c.Render(http.StatusOK, r.HTML("/department/new.plush.html"))
}

func Create(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	department := models.Department{}
	if err := c.Bind(&department); err != nil {
		return err
	}

	verrs, err := tx.Eager().ValidateAndCreate(&department)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("errors", verrs)
		requirements := models.RequirementTypes{}
		if err := tx.All(&requirements); err != nil {
			return err
		}

		if len(department.RequirementsType) == 0 {
			c.Flash().Add("danger", "Departments must have at least one type.")
		}

		c.Set("department", department)
		c.Set("requirements", requirements.Map())
		return c.Render(http.StatusUnprocessableEntity, r.HTML("/department/new.plush.html"))
	}

	for i := range department.RequirementsType {
		areaRequirementType := models.AreaRequirementType{}
		areaRequirementType.DepartmentID = department.ID
		areaRequirementType.RequirementTypeID = uuid.Must(uuid.FromString(i))
		err := tx.Create(&areaRequirementType)

		if err != nil {
			return err
		}
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

	return c.Render(http.StatusOK, r.HTML("/department/show.plush.html"))
}

func Edit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	department := models.Department{}
	departmentID := c.Param("department_id")

	requirementTypes := models.RequirementTypes{}
	if err := tx.All(&requirementTypes); err != nil {
		return err
	}

	areaRequirements := []models.AreaRequirementType{}
	if err := tx.Eager().All(&areaRequirements); err != nil {
		return err
	}

	err := tx.Eager("RequirementsTypes").Find(&department, departmentID)
	if err != nil {
		return err
	}

	c.Set("areaRequirements", areaRequirements)
	c.Set("requirements", requirementTypes.Map())
	c.Set("department", department)

	return c.Render(http.StatusOK, r.HTML("/department/edit.plush.html"))
}

func Update(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	department := models.Department{}
	departmentID := c.Param("department_id")

	err := tx.Eager("RequirementsTypes").Find(&department, departmentID)
	if err != nil {
		return err
	}

	if err := c.Bind(&department); err != nil {
		return err
	}

	tempAreaRequirementType := []models.AreaRequirementType{}
	if err := tx.Where("department_id = ?", departmentID).All(&tempAreaRequirementType); err != nil {
		return err
	}

	err = tx.Destroy(&tempAreaRequirementType)
	if err != nil {
		return err
	}

	for i := range department.RequirementsType {

		areaRequirementType := models.AreaRequirementType{}
		areaRequirementType.DepartmentID = department.ID
		areaRequirementType.RequirementTypeID = uuid.Must(uuid.FromString(i))

		err = tx.Save(&areaRequirementType)
		if err != nil {
			return err
		}
	}

	departments := []models.Department{}
	err = tx.All(&departments)
	if err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(&department)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("errors", verrs)
		requirements := models.RequirementTypes{}
		if err := tx.All(&requirements); err != nil {
			return err
		}

		if len(department.RequirementsType) == 0 {
			c.Flash().Add("danger", "Departments must have at least one type.")
		}

		c.Set("department", department)
		c.Set("requirements", requirements.Map())
		return c.Render(http.StatusUnprocessableEntity, r.HTML("/department/new.plush.html"))
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
