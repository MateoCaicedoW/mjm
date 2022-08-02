package requerimenttype_test

import (
	"mjm/app"
	"mjm/app/models"
	"net/http"
	"testing"

	"github.com/gobuffalo/suite/v4"
	"github.com/wawandco/fako"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	as := &ActionSuite{suite.NewAction(app.New())}
	suite.Run(t, as)
}

func (as *ActionSuite) Test_List() {

	department := models.Department{
		Name:        "name ",
		Description: "description",
	}

	as.NoError(as.DB.Create(&department))

	requirementType := models.RequirementType{}
	fako.Fill(&requirementType)
	requirementType.DepartmentID = department.ID
	as.NoError(as.DB.Create(&requirementType))

	res := as.HTML("/requirement-type/").Get()
	body := res.Body.String()
	as.Contains(body, requirementType.Name)

}
func (as *ActionSuite) Test_Create() {
	department := models.Department{
		Name:        "name ",
		Description: "description",
	}

	as.NoError(as.DB.Create(&department))

	requirementType := models.RequirementType{}
	requirementType.DepartmentID = department.ID
	fako.Fill(&requirementType)

	res := as.HTML("/requirement-type/create/").Post(requirementType)

	as.Equal(res.Code, http.StatusSeeOther)

	as.Equal("/requirement-type/", res.Location())

	RequirementTypes := models.RequirementTypes{}
	as.DB.All(&RequirementTypes)
	for _, v := range RequirementTypes {
		as.Equal(v.Name, requirementType.Name)
	}
}

func (as *ActionSuite) Test_Update() {
	department := models.Department{
		Name:        "name ",
		Description: "description",
	}

	as.NoError(as.DB.Create(&department))

	requirementType := &models.RequirementType{}
	requirementType.DepartmentID = department.ID
	fako.Fill(requirementType)
	as.NoError(as.DB.Create(requirementType))

	requirementTypeUpdate := &models.RequirementType{}
	fako.Fill(requirementTypeUpdate)
	requirementTypeUpdate.DepartmentID = department.ID

	res := as.HTML("/update/" + requirementType.ID.String()).Put(requirementTypeUpdate)
	as.Equal(res.Code, http.StatusSeeOther)
	as.Equal("/requirement-type/", res.Location())

}
func (as *ActionSuite) Test_Delete() {

	department := models.Department{
		Name:        "name ",
		Description: "description",
	}

	as.NoError(as.DB.Create(&department))

	requirementType := &models.RequirementType{}
	requirementType.DepartmentID = department.ID
	fako.Fill(requirementType)
	as.NoError(as.DB.Create(requirementType))

	res := as.HTML("/delete/" + requirementType.ID.String()).Delete()
	body := res.Body.String()
	as.NotContains(body, requirementType.Name)
	as.Equal(res.Code, http.StatusSeeOther)
	as.Equal("/requirement-type/", res.Location())

}
func (as *ActionSuite) Test_Edit() {

	department := models.Department{
		Name:        "name ",
		Description: "description",
	}

	as.NoError(as.DB.Create(&department))

	requirementType := &models.RequirementType{}
	requirementType.DepartmentID = department.ID
	fako.Fill(requirementType)
	as.NoError(as.DB.Create(requirementType))

	res := as.HTML("/edit/" + requirementType.ID.String()).Get()
	as.Equal(http.StatusOK, res.Code)
	body := res.Body.String()
	as.Contains(body, requirementType.Name)
	as.Contains(body, requirementType.ID.String())
	as.Contains(body, "Edit RequirementType")

}

func (as *ActionSuite) Test_Show() {
	department := models.Department{
		Name:        "name ",
		Description: "description",
	}

	as.NoError(as.DB.Create(&department))

	requirementType := &models.RequirementType{}
	fako.Fill(requirementType)
	requirementType.DepartmentID = department.ID
	as.NoError(as.DB.Create(requirementType))

	res := as.HTML("/Show/" + requirementType.ID.String()).Get()
	as.Equal(http.StatusOK, res.Code)
	body := res.Body.String()
	as.Contains(body, requirementType.Name)
	as.Contains(body, requirementType.ID.String())

}

func (as *ActionSuite) Test_New() {
	res := as.HTML("/requirement-type/new/").Get()
	as.Equal(http.StatusOK, res.Code)
	body := res.Body.String()

	as.Contains(body, "New RequirementType")
	as.Contains(body, "Save")
}
