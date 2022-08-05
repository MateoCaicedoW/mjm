package app

import (
	"net/http"

	"mjm/app/actions/departments"
	"mjm/app/actions/home"
	requirement_type "mjm/app/actions/requeriment_type"
	"mjm/app/actions/users"

	"mjm/app/middleware"
	"mjm/public"

	"github.com/gobuffalo/buffalo"
)

func setRoutes(root *buffalo.App) {
	root.Use(middleware.RequestID)
	root.Use(middleware.Database)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/", home.Index)
	requirementType := root.Group("/requirement-types")
	requirementType.GET("/", requirement_type.List)
	requirementType.GET("/new", requirement_type.New)
	requirementType.POST("/create", requirement_type.Create)
	requirementType.GET("/show/{requirement_type_id}", requirement_type.Show).Name("showRequirementTypes")
	requirementType.GET("/edit/{requirement_type_id}", requirement_type.Edit).Name("editRequirementTypes")
	requirementType.PUT("/update/{requirement_type_id}", requirement_type.Update).Name("updateRequirementTypes")
	requirementType.DELETE("/delete/{requirement_type_id}", requirement_type.Delete).Name("deleteRequirementTypes")

	department := root.Group("/departments")
	department.GET("/", departments.List)
	department.GET("/new", departments.New)
	department.POST("/create", departments.Create)
	department.GET("/edit/{department_id}", departments.Edit).Name("editDepartments")
	department.GET("/show/{department_id}", departments.Show).Name("showDepartments")
	department.PUT("/update/{department_id}", departments.Update).Name("updateDepartments")
	department.DELETE("/destroy/{department_id}", departments.Destroy).Name("destroyDepartments")

	user := root.Group("/users")
	user.GET("/", users.List)
	user.GET("/new", users.New)
	user.POST("/create", users.Create)
	user.GET("/edit/{user_id}", users.Edit).Name("editUsers")
	user.PUT("/update/{user_id}", users.Update).Name("updateUsers")
	user.DELETE("/delete/{user_id}", users.Delete).Name("deleteUsers")
	user.GET("/view/{user_id}", users.View).Name("viewUsers")

	root.ServeFiles("/", http.FS(public.FS()))
}
