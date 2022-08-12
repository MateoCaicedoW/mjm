package requirements_test

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
	requirements := [2]models.Requirement{}
	department, user, requirementType, requirementSubType := Create(*as)

	for i := 0; i < len(requirements); i++ {
		requirements[i].RequirementTypeID = requirementType.ID
		requirements[i].RequirementSubTypeID = requirementSubType.ID
		requirements[i].CreatedByUserID = user.ID
		requirements[i].RequestingDepartmentID = department.ID
		requirements[i].ServiceDepartmentID = department.ID

		fako.Fill(&requirements[i])
		err5 := as.DB.Create(&requirements[i], "modified_by", "approved_by", "declined_by", "accepted_by", "finished_by", "proccessed_by", "assigned_to", "assigned_by")
		as.NoError(err5)
		res := as.HTML("/requirements").Get()
		body := res.Body.String()
		as.Contains(body, requirements[i].Title)
	}

	as.DB.Reload(&requirements)
	as.Equal(len(requirements), 2)

}

func (as *ActionSuite) Test_Create() {
	requirement := models.Requirement{}
	requirementArray := []models.Requirement{}
	department, user, requirementType, requirementSubType := Create(*as)

	requirement.RequirementTypeID = requirementType.ID
	requirement.RequirementSubTypeID = requirementSubType.ID
	requirement.CreatedByUserID = user.ID
	requirement.RequestingDepartmentID = department.ID
	requirement.ServiceDepartmentID = department.ID

	fako.Fill(&requirement)

	req := as.HTML("/requirements/create").Post(requirement)

	req2 := as.HTML("/requirements").Get()

	body := req2.Body.String()

	as.Contains(body, requirement.Title)
	as.Contains(body, requirement.Description)

	as.Equal(http.StatusSeeOther, req.Code)
	as.Equal("/requirements", req.Location())
	as.DB.Reload(&requirement)
	as.DB.All(&requirementArray)
	as.Equal(1, len(requirementArray))

}

func (as *ActionSuite) Test_Create_Failed() {
	requirements := models.Requirement{}

	department, _, requirementType, requirementSubType := Create(*as)

	requirements.RequirementTypeID = requirementType.ID
	requirements.RequirementSubTypeID = requirementSubType.ID

	requirements.RequestingDepartmentID = department.ID
	requirements.ServiceDepartmentID = department.ID

	fako.Fill(&requirements)
	req := as.HTML("/requirements/create").Post(requirements)
	as.Equal(http.StatusInternalServerError, req.Code)
}

func (as *ActionSuite) Test_New() {
	res := as.HTML("/requirements/new").Get()
	as.Equal(http.StatusOK, res.Code)

	body := res.Body.String()

	as.Contains(body, "New Requirement")

}

func (as *ActionSuite) Test_Edit() {
	requirements := models.Requirement{}

	department, user, requirementType, requirementSubType := Create(*as)
	requirements.RequirementTypeID = requirementType.ID
	requirements.RequirementSubTypeID = requirementSubType.ID
	requirements.CreatedByUserID = user.ID
	requirements.RequestingDepartmentID = department.ID
	requirements.ServiceDepartmentID = department.ID

	fako.Fill(&requirements)
	err := as.DB.Create(&requirements)
	as.NoError(err)
	res := as.HTML("/requirements/edit/" + requirements.ID.String()).Get()
	as.Equal(http.StatusOK, res.Code)
	body := res.Body.String()
	as.Contains(body, "Edit Requirement")
	as.NotContains(body, "New Requirement")

}

func (as *ActionSuite) Test_Update() {
	requirements := &models.Requirement{}

	department, user, requirementType, requirementSubType := Create(*as)

	requirements.RequirementTypeID = requirementType.ID
	requirements.RequirementSubTypeID = requirementSubType.ID
	requirements.CreatedByUserID = user.ID
	requirements.RequestingDepartmentID = department.ID
	requirements.ServiceDepartmentID = department.ID

	fako.Fill(requirements)
	err := as.DB.Create(requirements)
	as.NoError(err)

	requirementsUpdate := &models.Requirement{}
	fako.Fill(requirementsUpdate)
	requirementsUpdate.ID = requirements.ID
	requirementsUpdate.RequirementTypeID = requirementType.ID
	requirementsUpdate.RequirementSubTypeID = requirementSubType.ID
	requirementsUpdate.CreatedByUserID = user.ID
	requirementsUpdate.RequestingDepartmentID = department.ID
	requirementsUpdate.ServiceDepartmentID = department.ID

	res := as.HTML("/requirements/update/" + requirements.ID.String()).Put(requirementsUpdate)

	as.Equal(http.StatusSeeOther, res.Code)
	as.Equal("/requirements", res.Location())
	as.DB.Reload(requirements)
	as.Equal(requirementsUpdate.Title, requirements.Title)

}

func (as *ActionSuite) Test_Destroy() {
	requirements := &models.Requirement{}

	department, user, requirementType, requirementSubType := Create(*as)
	requirements.RequirementTypeID = requirementType.ID
	requirements.RequirementSubTypeID = requirementSubType.ID
	requirements.CreatedByUserID = user.ID
	requirements.RequestingDepartmentID = department.ID
	requirements.ServiceDepartmentID = department.ID

	fako.Fill(requirements)
	err := as.DB.Create(requirements)
	as.NoError(err)

	res := as.HTML("/requirements/delete/" + requirements.ID.String()).Delete()

	as.Equal(http.StatusSeeOther, res.Code)
	as.Equal("/requirements", res.Location())
	as.DB.Reload(requirements)
	count, _ := as.DB.Count(requirements)
	as.Equal(0, count)

}

func Create(as ActionSuite) (models.Department, models.User, models.RequirementType, models.RequirementSubType) {
	deparment := models.Department{}
	fako.Fill(&deparment)
	err := as.DB.Create(&deparment)
	as.NoError(err)

	user := models.User{}
	user.DepartmentID = deparment.ID
	fako.Fill(&user)
	err2 := as.DB.Create(&user)
	as.NoError(err2)

	requirementType := models.RequirementType{}
	fako.Fill(&requirementType)
	err3 := as.DB.Create(&requirementType)
	as.NoError(err3)

	requirementSubType := models.RequirementSubType{}
	requirementSubType.RequirementTypeID = requirementType.ID
	fako.Fill(&requirementSubType)
	err4 := as.DB.Create(&requirementSubType)
	as.NoError(err4)

	return deparment, user, requirementType, requirementSubType
}