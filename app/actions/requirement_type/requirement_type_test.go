package requirement_type_test

import (
	"fmt"
	"mjm/app"
	"mjm/app/models"
	"net/http"
	"testing"

	"github.com/gobuffalo/suite/v4"
	"github.com/gofrs/uuid"
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

	user := models.User{
		FirstName: "jhon",
		LastName:  "Smith",
	}
	department := models.Department{
		Name:        "name ",
		Description: "description",
	}
	as.NoError(as.DB.Create(&department))
	user.DepartmentID = department.ID
	as.NoError(as.DB.Create(&user))

	requirementType := models.RequirementType{}
	fako.Fill(&requirementType)
	requirementType.DepartmentID = department.ID
	requirementType.CreatedByUserID = user.ID
	as.NoError(as.DB.Create(&requirementType))

	res := as.HTML("/requirement-types/").Get()
	body := res.Body.String()
	as.Contains(body, requirementType.Name)

}

func (as *ActionSuite) Test_Create() {
	user := models.User{
		ID:        uuid.Must(uuid.FromString("175afda1-82ef-4950-b8db-6dab15740d63")),
		FirstName: "jhon",
		LastName:  "Smith",
	}
	department := models.Department{
		Name:        "name ",
		Description: "description",
	}
	as.NoError(as.DB.Create(&department))
	user.DepartmentID = department.ID
	as.NoError(as.DB.Create(&user))

	requirementType := &models.RequirementType{}
	fako.Fill(requirementType)
	requirementType.CreatedByUserID = user.ID
	requirementType.DepartmentID = department.ID

	fmt.Println(requirementType)
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
	user := models.User{
		ID:        uuid.Must(uuid.FromString("175afda1-82ef-4950-b8db-6dab15740d63")),
		FirstName: "jhon",
		LastName:  "Smith",
	}
	department := models.Department{
		Name:        "name ",
		Description: "description",
	}
	as.NoError(as.DB.Create(&department))
	user.DepartmentID = department.ID
	as.NoError(as.DB.Create(&user))

	requirementType := &models.RequirementType{}
	requirementType.DepartmentID = department.ID
	requirementType.CreatedByUserID = user.ID
	fako.Fill(requirementType)
	as.NoError(as.DB.Create(requirementType))

	requirementTypeUpdate := &models.RequirementType{}
	fako.Fill(requirementTypeUpdate)
	requirementTypeUpdate.CreatedByUserID = user.ID
	requirementTypeUpdate.DepartmentID = department.ID
	requirementTypeUpdate.ID = requirementType.ID

	res := as.HTML("/requirement-types/update/%v", requirementTypeUpdate.ID).Put(requirementTypeUpdate)
	as.Equal(res.Code, http.StatusSeeOther)
	as.Equal("/requirement-types/", res.Location())
	as.DB.Reload(requirementType)
	as.Equal(requirementTypeUpdate.Name, requirementType.Name)

}
func (as *ActionSuite) Test_Delete() {
	user := models.User{
		ID:        uuid.Must(uuid.FromString("175afda1-82ef-4950-b8db-6dab15740d63")),
		FirstName: "jhon",
		LastName:  "Smith",
	}
	department := models.Department{
		Name:        "name ",
		Description: "description",
	}
	as.NoError(as.DB.Create(&department))
	user.DepartmentID = department.ID
	as.NoError(as.DB.Create(&user))

	requirementType := &models.RequirementType{}
	requirementType.CreatedByUserID = user.ID
	requirementType.DepartmentID = department.ID
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
	user := models.User{
		ID:        uuid.Must(uuid.FromString("175afda1-82ef-4950-b8db-6dab15740d63")),
		FirstName: "jhon",
		LastName:  "Smith",
	}
	department := models.Department{
		Name:        "name ",
		Description: "description",
	}
	as.NoError(as.DB.Create(&department))
	user.DepartmentID = department.ID
	as.NoError(as.DB.Create(&user))

	requirementType := &models.RequirementType{}
	requirementType.DepartmentID = department.ID
	requirementType.CreatedByUserID = user.ID
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
	user := models.User{
		ID:        uuid.Must(uuid.FromString("175afda1-82ef-4950-b8db-6dab15740d63")),
		FirstName: "jhon",
		LastName:  "Smith",
	}
	department := models.Department{
		Name:        "name ",
		Description: "description",
	}
	as.NoError(as.DB.Create(&department))
	user.DepartmentID = department.ID
	as.NoError(as.DB.Create(&user))

	requirementType := &models.RequirementType{}
	fako.Fill(requirementType)
	requirementType.DepartmentID = department.ID
	requirementType.CreatedByUserID = user.ID
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
