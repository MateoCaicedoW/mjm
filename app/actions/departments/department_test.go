package departments_test

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

func (as ActionSuite) Test_ListDepartment() {
	deparments := models.Department{
		Name:        "Sistemas",
		Description: "Description",
	}

	err := as.DB.Create(&deparments)
	as.NoError(err)
	res := as.HTML("/departments").Get()
	body := res.Body.String()
	as.Contains(body, deparments.Name)
	as.Contains(body, deparments.Description)
}

func (as ActionSuite) Test_CreateDepartment() {
	deparments := models.Department{
		Name:        "Sistemas",
		Description: "Description",
	}

	res := as.HTML("/create-department").Post(deparments)

	as.Equal(res.Code, http.StatusSeeOther)
	as.Equal("/departments", res.Location())
}

func (as ActionSuite) Test_UpdateDepartment() {
	deparments := &models.Department{}
	fako.Fill(deparments)
	err := as.DB.Create(deparments)
	as.NoError(err)

	deparmentsUpdate := &models.Department{}
	fako.Fill(deparmentsUpdate)

	res := as.HTML("/edit-data/%s", deparments.ID).Put(deparmentsUpdate)

	as.Equal(res.Code, http.StatusSeeOther)
	as.Equal("/departments", res.Location())
	as.DB.Reload(deparments)
	as.Equal(deparmentsUpdate.Name, deparments.Name)
}

func (as ActionSuite) Test_DestroyDepartment() {
	deparments := &models.Department{}
	fako.Fill(deparments)
	err := as.DB.Create(deparments)
	as.NoError(err)

	res := as.HTML("/delete/%s", deparments.ID).Delete()

	as.Equal(res.Code, http.StatusSeeOther)
	as.Equal("/departments", res.Location())
}

func (as ActionSuite) Test_NewDepartment() {
	res := as.HTML("/add-department").Get()

	as.Equal(res.Code, http.StatusOK)
	body := res.Body.String()
	as.Contains(body, "Save Line")
}

func (as *ActionSuite) Test_EditDepartment() {
	deparments := &models.Department{}
	fako.Fill(deparments)
	err := as.DB.Create(deparments)
	as.NoError(err)

	res := as.HTML("/edit/" + deparments.ID.String()).Get()
	as.Equal(http.StatusOK, res.Code)

	body := res.Body.String()
	as.Contains(body, deparments.Name)
	as.Contains(body, "Save Line")
}

func (as *ActionSuite) Test_ViewDepartment() {
	deparments := &models.Department{}
	fako.Fill(deparments)
	err := as.DB.Create(deparments)
	as.NoError(err)

	res := as.HTML("/view/" + deparments.ID.String()).Get()
	as.Equal(http.StatusOK, res.Code)

	body := res.Body.String()
	as.Contains(body, deparments.Description)
}