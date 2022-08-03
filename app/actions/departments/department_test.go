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

func (as ActionSuite) Test_List() {
	deparment := models.Department{}
	fako.Fill(&deparment)

	err := as.DB.Create(&deparment)
	as.NoError(err)
	res := as.HTML("/departments/list").Get()
	body := res.Body.String()
	as.Contains(body, deparment.Name)
	as.Contains(body, deparment.Description)
}

func (as *ActionSuite) Test_Create() {
	department := &models.Department{}
	fako.Fill(department)

	count, _ := as.DB.Count(&department)

	res := as.HTML("/department/create/").Post(department)

	as.Equal(res.Code, http.StatusSeeOther)
	as.Equal("/departments/list", res.Location())

	departmentCreate := []models.Department{}
	as.DB.All(&departmentCreate)

	countDB, _ := as.DB.Count(&department)

	for _, v := range departmentCreate {
		as.Equal(v.Name, department.Name)
	}

	as.NotEqual(count, countDB)
}

func (as *ActionSuite) Test_Update() {
	deparment := &models.Department{}
	fako.Fill(deparment)
	err := as.DB.Create(deparment)
	as.NoError(err)

	departmentUpdate := &models.Department{}
	fako.Fill(departmentUpdate)
	departmentUpdate.ID = deparment.ID

	res := as.HTML("/update/%s", departmentUpdate.ID).Put(departmentUpdate)

	as.Equal(res.Code, http.StatusSeeOther)
	as.Equal("/departments/list", res.Location())
	as.DB.Reload(deparment)
	as.Equal(deparment.Name, departmentUpdate.Name)
}

func (as *ActionSuite) Test_Destroy() {
	deparment := &models.Department{}
	fako.Fill(deparment)
	err := as.DB.Create(deparment)
	as.NoError(err)

	res := as.HTML("/destroy/%s", deparment.ID).Delete()

	as.Equal(res.Code, http.StatusSeeOther)
	as.Equal("/departments/list", res.Location())

	body := res.Body.String()
	as.NotContains(body, deparment.Name)
	as.NotContains(body, deparment.Description)
}

func (as *ActionSuite) Test_New() {
	res := as.HTML("/department/new").Get()

	as.Equal(res.Code, http.StatusOK)
	body := res.Body.String()
	as.Contains(body, "Save Line")
}

func (as *ActionSuite) Test_Edit() {
	deparment := &models.Department{}
	fako.Fill(deparment)
	err := as.DB.Create(deparment)
	as.NoError(err)

	res := as.HTML("/edit/" + deparment.ID.String()).Get()
	as.Equal(http.StatusOK, res.Code)

	body := res.Body.String()
	as.Contains(body, deparment.Name)
	as.Contains(body, "Save Line")
}

func (as *ActionSuite) Test_View() {
	deparment := &models.Department{}
	fako.Fill(deparment)
	err := as.DB.Create(deparment)
	as.NoError(err)

	res := as.HTML("/show/" + deparment.ID.String()).Get()
	as.Equal(http.StatusOK, res.Code)

	body := res.Body.String()
	as.Contains(body, deparment.Description)
}
