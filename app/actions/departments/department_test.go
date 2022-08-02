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
	deparments := models.Department{}
	fako.Fill(&deparments)

	err := as.DB.Create(&deparments)
	as.NoError(err)
	res := as.HTML("/departments/list").Get()
	body := res.Body.String()
	as.Contains(body, deparments.Name)
	as.Contains(body, deparments.Description)
}

func (as ActionSuite) Test_Create() {
	deparments := &models.Department{}
	fako.Fill(deparments)

	res := as.HTML("/department/create").Post(deparments)

	as.Equal(res.Code, http.StatusSeeOther)
	as.Equal("/departments/list", res.Location())
}

func (as ActionSuite) Test_Update() {
	deparments := &models.Department{}
	fako.Fill(deparments)
	err := as.DB.Create(deparments)
	as.NoError(err)

	departmentsUpdate := &models.Department{}
	fako.Fill(departmentsUpdate)
	departmentsUpdate.ID = deparments.ID

	res := as.HTML("/update/%s", departmentsUpdate.ID).Put(departmentsUpdate)

	as.Equal(res.Code, http.StatusSeeOther)
	as.Equal("/departments/list", res.Location())
	as.DB.Reload(deparments)
	as.Equal(departmentsUpdate.Name, deparments.Name)
}

func (as ActionSuite) Test_Destroy() {
	deparments := &models.Department{}
	fako.Fill(deparments)
	err := as.DB.Create(deparments)
	as.NoError(err)

	res := as.HTML("/destroy/%s", deparments.ID).Delete()

	as.Equal(res.Code, http.StatusSeeOther)
	as.Equal("/departments/list", res.Location())
}

func (as ActionSuite) Test_New() {
	res := as.HTML("/department/new").Get()

	as.Equal(res.Code, http.StatusOK)
	body := res.Body.String()
	as.Contains(body, "Save Line")
}

func (as *ActionSuite) Test_Edit() {
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

func (as *ActionSuite) Test_View() {
	deparments := &models.Department{}
	fako.Fill(deparments)
	err := as.DB.Create(deparments)
	as.NoError(err)

	res := as.HTML("/show/" + deparments.ID.String()).Get()
	as.Equal(http.StatusOK, res.Code)

	body := res.Body.String()
	as.Contains(body, deparments.Description)
}
