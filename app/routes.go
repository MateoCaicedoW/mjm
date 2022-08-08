package app

import (
	"net/http"

	"mjm/app/actions/departments"
	"mjm/app/actions/requirements"

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

	requirementType := root.Group("/requirement-types")
	requirementType.GET("/", requirement_type.List)
	requirementType.GET("/new", requirement_type.New).Name("newRequirementType")
	requirementType.POST("/create", requirement_type.Create).Name("createRequirementType")
	requirementType.GET("/show/{requirement_type_id}", requirement_type.Show).Name("showRequirementType")
	requirementType.GET("/edit/{requirement_type_id}", requirement_type.Edit).Name("editRequirementType")
	requirementType.PUT("/update/{requirement_type_id}", requirement_type.Update).Name("updateRequirementType")
	requirementType.DELETE("/delete/{requirement_type_id}", requirement_type.Delete).Name("deleteRequirementType")

	department := root.Group("/departments")
	department.GET("/", departments.List)
	department.GET("/new", departments.New).Name("newDepartment")
	department.POST("/create", departments.Create).Name("createDepartment")
	department.GET("/edit/{department_id}", departments.Edit).Name("editDepartment")
	department.GET("/show/{department_id}", departments.Show).Name("showDepartment")
	department.PUT("/update/{department_id}", departments.Update).Name("updateDepartment")
	department.DELETE("/delete/{department_id}", departments.Delete).Name("deleteDepartment")

	user := root.Group("/users")
	user.GET("/", users.List)
	user.GET("/new", users.New).Name("newUser")
	user.POST("/create", users.Create).Name("createUser")
	user.GET("/show/{user_id}", users.View).Name("showUser")
	user.GET("/edit/{user_id}", users.Edit).Name("editUser")
	user.PUT("/update/{user_id}", users.Update).Name("updateUser")
	user.DELETE("/delete/{user_id}", users.Delete).Name("deleteUser")

	requirement := root.Group("/requirements")
	requirement.GET("/", requirements.List)
	requirement.GET("/new", requirements.New).Name("newRequirement")
	requirement.POST("/create", requirements.Create).Name("createRequirement")
	requirement.GET("/show/{requirement_id}", requirements.Show).Name("showRequirement")
	requirement.GET("/edit/{requirement_id}", requirements.Edit).Name("editRequirement")
	requirement.PUT("/update/{requirement_id}", requirements.Update).Name("updateRequirement")
	requirement.DELETE("/delete/{requirement_id}", requirements.Destroy).Name("deleteRequirement")

	root.ServeFiles("/", http.FS(public.FS()))
}
