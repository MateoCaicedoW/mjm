package requirements_test

import (
	"mjm/app"
	"mjm/app/models"
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
	requirementType.DepartmentID = deparment.ID
	fako.Fill(&requirementType)
	err3 := as.DB.Create(&requirementType)
	as.NoError(err3)

	requirementSubType := models.RequirementSubType{}
	requirementSubType.RequirementTypeID = requirementType.ID
	fako.Fill(&requirementSubType)
	err4 := as.DB.Create(&requirementSubType)
	as.NoError(err4)

	for i := 0; i < len(requirements); i++ {
		requirements[i].RequirementTypeID = requirementType.ID
		requirements[i].RequirementSubTypeID = requirementSubType.ID
		requirements[i].CreatedByUserID = user.ID
		requirements[i].RequestingDepartmentID = deparment.ID
		requirements[i].ServiceDepartmentID = deparment.ID
		fako.Fill(&requirements[i])
		err5 := as.DB.Create(&requirements[i])
		as.NoError(err5)
	}

}
