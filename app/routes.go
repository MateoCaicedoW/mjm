package app

import (
	"net/http"

	"mjm/app/actions/departments"
	"mjm/app/actions/home"
	requirement_type "mjm/app/actions/requeriment_type"
	"mjm/app/actions/user"

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
	root.GET("/requirement-types", requirement_type.List)
	root.GET("/requirement-type/new", requirement_type.New)
	root.POST("/requirement-type/create", requirement_type.Create)
	root.GET("/show/{requirement_type_id}", requirement_type.Show)
	root.GET("/edit/{requirement_type_id}", requirement_type.Edit)
	root.PUT("/update/{requirement_type_id}", requirement_type.Update)
	root.DELETE("/delete/{requirement_type_id}", requirement_type.Delete)
  
	root.GET("/departments/list", departments.List)
	root.GET("/department/new", departments.New)
	root.POST("/department/create", departments.Create)
	root.GET("/edit/{department_id}", departments.Edit)
	root.GET("/show/{department_id}", departments.Show)
	root.PUT("/update/{department_id}", departments.Update)
	root.DELETE("/destroy/{department_id}", departments.Destroy)

	root.GET("/users", user.List)
	root.GET("/new-users", user.New)
	root.POST("/create-user", user.Create)
	root.GET("/edit-user/{user_id}", user.Edit)
	root.PUT("/update-user/{user_id}", user.Update)
	root.DELETE("/delete-user/{user_id}", user.Delete)
	root.GET("/view-user/{user_id}", user.View)


	root.ServeFiles("/", http.FS(public.FS()))
}
