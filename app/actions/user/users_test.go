package user_test

import (
	"mjm/app"
	"mjm/app/models"
	"net/http"
	"testing"

	"github.com/gobuffalo/suite/v4"
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

	user := models.User{
		FirstName:    "Joaquin",
		LastName:     "Olivo",
		PhoneNumber:  "3042015706",
		DepartmentID: department.ID,
	}

	as.NoError(as.DB.Create(&user))

	res := as.HTML("/users").Get()
	body := res.Body.String()
	as.Contains(body, user.FirstName)

}

func (as *ActionSuite) Test_Create() {

	department := models.Department{
		Name:        "name ",
		Description: "description",
	}

	as.NoError(as.DB.Create(&department))

	user := models.User{
		FirstName:    "Joaquin",
		LastName:     "Olivo",
		PhoneNumber:  "3042015706",
		DepartmentID: department.ID,
	}

	res := as.HTML("/create-user").Post(user)

	as.Equal(res.Code, http.StatusSeeOther)
	as.Equal("/users/", res.Location())
	userArray := []models.User{}
	as.DB.All(userArray)
	for _, v := range userArray {
		as.Equal(v.FirstName, user.FirstName)
	}

}

func (as *ActionSuite) Test_Update() {
	department := models.Department{
		Name:        "name ",
		Description: "description",
	}

	as.NoError(as.DB.Create(&department))

	user := models.User{
		FirstName:    "Joaquin",
		LastName:     "Olivo",
		PhoneNumber:  "3042015706",
		DepartmentID: department.ID,
	}
	as.NoError(as.DB.Create(&user))

	userUpdate := &models.User{}
	userUpdate.DNI = "1k123j1j43203k4"
	userUpdate.EmailAddress = "jolivo@wawand.co"
	userUpdate.PhoneNumber = "321234543"
	userUpdate.FirstName = "Joaquin"
	userUpdate.LastName = "asfasf"
	userUpdate.ID = user.ID
	userUpdate.DepartmentID = department.ID

	res := as.HTML("/update-user/" + user.ID.String()).Put(userUpdate)
	as.Equal(res.Code, http.StatusSeeOther)
	as.Equal("/users/", res.Location())
	as.DB.Reload(&user)
	as.Equal(userUpdate.FirstName, user.FirstName)

}

func (as *ActionSuite) Test_Delete() {
	department := models.Department{
		Name:        "name ",
		Description: "description",
	}

	as.NoError(as.DB.Create(&department))

	user := models.User{
		FirstName:    "Joaquin",
		LastName:     "Olivo",
		PhoneNumber:  "3042015706",
		DepartmentID: department.ID,
	}
	as.NoError(as.DB.Create(&user))

	userUpdate := &models.User{}
	userUpdate.DNI = "1k123j1j43203k4"
	userUpdate.EmailAddress = "jolivo@wawand.co"
	userUpdate.PhoneNumber = "321234543"
	userUpdate.FirstName = "Joaquin"
	userUpdate.LastName = "asfasf"
	userUpdate.ID = user.ID
	userUpdate.DepartmentID = department.ID

	res := as.HTML("/delete-user/" + user.ID.String()).Delete()
	as.Equal(res.Code, http.StatusSeeOther)
	as.Equal("/users/", res.Location())

}

func (as *ActionSuite) Test_Edit() {

	department := models.Department{
		Name:        "name ",
		Description: "description",
	}

	as.NoError(as.DB.Create(&department))

	user := &models.User{
		FirstName:    "Joaquin",
		LastName:     "Olivo",
		PhoneNumber:  "3042015706",
		DepartmentID: department.ID,
	}

	as.NoError(as.DB.Create(user))

	res := as.HTML("/users/").Get()
	as.Equal(http.StatusOK, res.Code)
	body := res.Body.String()
	as.Contains(body, user.FirstName)

}

func (as *ActionSuite) Test_Form_Create() {

	user := &models.User{}

	res := as.HTML("/new-users/").Get()
	as.Equal(http.StatusOK, res.Code)
	body := res.Body.String()
	as.Contains(body, user.FirstName)
	as.Contains(body, user.LastName)
	as.Contains(body, user.PhoneNumber)
	as.Contains(body, user.DNI)
	as.Contains(body, "New User")

}

func (as *ActionSuite) Test_Info_User() {

	department := models.Department{
		Name:        "name ",
		Description: "description",
	}

	as.NoError(as.DB.Create(&department))

	user := &models.User{
		FirstName:    "Joaquin",
		LastName:     "Olivo",
		PhoneNumber:  "3042015706",
		DepartmentID: department.ID,
	}

	as.NoError(as.DB.Create(user))

	res := as.HTML("/view-user/" + user.ID.String()).Get()
	as.Equal(http.StatusOK, res.Code)
	body := res.Body.String()
	as.Contains(body, user.FirstName)
	as.Contains(body, user.ID.String())

}
