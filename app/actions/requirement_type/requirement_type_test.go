package requirement_type_test

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
	requirementType := models.RequirementType{}
	fako.Fill(&requirementType)
	as.NoError(as.DB.Create(&requirementType))

	res := as.HTML("/requirement-types/").Get()
	body := res.Body.String()
	as.Contains(body, requirementType.Name)

}

func (as *ActionSuite) Test_Create() {
	requirementType := models.RequirementType{}
	fako.Fill(&requirementType)

	res := as.HTML("/requirement-types/create/").Post(requirementType)
	as.Equal(http.StatusSeeOther, res.Code)
	as.Equal("/requirement-types/", res.Location())

	requirementTypes := models.RequirementTypes{}
	as.DB.All(&requirementTypes)
	for _, v := range requirementTypes {
		as.Equal(v.Name, requirementType.Name)
	}
}

func (as *ActionSuite) Test_Update() {
	requirementType := &models.RequirementType{}
	fako.Fill(requirementType)
	requirementType.Description = "Descriptions"
	as.NoError(as.DB.Create(requirementType))

	requirementTypeUpdate := &models.RequirementType{}
	fako.Fill(requirementTypeUpdate)
	requirementTypeUpdate.ID = requirementType.ID
	requirementTypeUpdate.Description = requirementType.Description

	res := as.HTML("/requirement-types/update/%v", requirementTypeUpdate.ID).Put(requirementTypeUpdate)
	as.Equal(res.Code, http.StatusSeeOther)
	as.Equal("/requirement-types/", res.Location())
	as.DB.Reload(requirementType)
	as.Equal(requirementTypeUpdate.Name, requirementType.Name)

}

func (as *ActionSuite) Test_Delete() {
	requirementType := &models.RequirementType{}
	fako.Fill(requirementType)
	as.NoError(as.DB.Create(requirementType))

	res := as.HTML("/requirement-types/delete/" + requirementType.ID.String()).Delete()
	as.Equal("/requirement-types/", res.Location())
	as.Equal(res.Code, http.StatusSeeOther)

	requirementTypes := models.RequirementTypes{}
	as.DB.All(&requirementTypes)
	count, _ := as.DB.Count(requirementTypes)
	as.Equal(count, 0)
	res = as.HTML("/requirement-types/").Get()
	body := res.Body.String()
	as.NotContains(body, requirementType.Name)

}

func (as *ActionSuite) Test_Edit() {
	requirementType := &models.RequirementType{}
	fako.Fill(requirementType)
	as.NoError(as.DB.Create(requirementType))

	res := as.HTML("/requirement-types/edit/" + requirementType.ID.String()).Get()
	as.Equal(http.StatusOK, res.Code)
	body := res.Body.String()
	as.Contains(body, requirementType.Name)
	as.Contains(body, requirementType.ID.String())
	as.Contains(body, "Edit RequirementType")
}

func (as *ActionSuite) Test_Show() {
	requirementType := &models.RequirementType{}
	fako.Fill(requirementType)
	as.NoError(as.DB.Create(requirementType))

	res := as.HTML("/requirement-types/show/" + requirementType.ID.String()).Get()
	as.Equal(http.StatusOK, res.Code)
	body := res.Body.String()
	as.Contains(body, requirementType.Name)
	as.Contains(body, requirementType.ID.String())
}

func (as *ActionSuite) Test_New() {
	res := as.HTML("/requirement-types/new/").Get()
	as.Equal(http.StatusOK, res.Code)
	body := res.Body.String()

	as.Contains(body, "New RequirementType")
	as.Contains(body, "Save")
}
