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

	department, user, requirementType, requirementSubType := Create(*as)

	test := []struct {
		requirement          []models.Requirement
		elementsContained    []string
		elementNotContained  []string
		numberOfRequirements int
	}{

		{
			requirement: []models.Requirement{
				{
					Title:       "Test",
					Description: "Test",
				},
			},
			elementsContained:    []string{"Test"},
			elementNotContained:  []string{},
			numberOfRequirements: 1,
		},

		{
			requirement: []models.Requirement{
				{
					Title:       "Test1",
					Description: "Test",
				},
				{
					Title:       "Test2",
					Description: "Test",
				},
				{
					Title:       "Test3",
					Description: "Test",
				},
			},
			elementsContained:    []string{"Test1", "Test2", "Test3"},
			elementNotContained:  []string{},
			numberOfRequirements: 3,
		},

		{
			requirement: []models.Requirement{
				{
					Title:       "",
					Description: "",
				},
			},
			elementsContained:    []string{},
			elementNotContained:  []string{"Test"},
			numberOfRequirements: 0,
		},
	}

	for _, t := range test {
		as.DB.Reload(&t.requirement)
		for _, r := range t.requirement {
			if t.numberOfRequirements == 0 {
				as.Error(as.DB.Create(&r))
			}
			if r.Title != "" {
				r.RequirementTypeID = requirementType.ID
				r.RequirementSubTypeID = requirementSubType.ID
				r.CreatedByUserID = user.ID
				r.RequestingDepartmentID = department.ID
				r.ServiceDepartmentID = department.ID
				as.NoError(as.DB.Create(&r))
			}

		}

		requirements := []models.Requirement{}
		as.NoError(as.DB.All(&requirements))

		as.Equal(t.numberOfRequirements, len(requirements))

		res := as.HTML("/requirements").Get()
		as.Equal(http.StatusOK, res.Code)
		body := res.Body.String()

		if len(t.elementsContained) > 0 {
			for _, e := range t.elementsContained {
				as.Contains(body, e)
			}

		}
		if len(t.elementNotContained) > 0 {
			for _, e := range t.elementNotContained {
				as.NotContains(body, e)
			}

		}

		as.NoError(as.DB.Destroy(&requirements))

	}

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

	user := models.User{
		DepartmentID: deparment.ID,
		FirstName:    "John",
		LastName:     "Doe",
		EmailAddress: "johndoe@wawand.co",
		DNI:          "12345678",
		PhoneNumber:  "12345678",
	}

	err2 := as.DB.Create(&user)
	as.NoError(err2)

	requirementType := models.RequirementType{
		Name:         "Requirement Type",
		DepartmentID: deparment.ID,
	}

	err3 := as.DB.Create(&requirementType)
	as.NoError(err3)

	requirementSubType := models.RequirementSubType{
		Name:              "Requirement Sub Type",
		RequirementTypeID: requirementType.ID,
	}

	err4 := as.DB.Create(&requirementSubType)
	as.NoError(err4)

	return deparment, user, requirementType, requirementSubType
}