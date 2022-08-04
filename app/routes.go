package app

import (
	"net/http"

	"mjm/app/actions/departments"
	"mjm/app/actions/home"
	"mjm/app/middleware"
	"mjm/public"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.RequestID)
	root.Use(middleware.Database)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/", home.Index)
	root.GET("/departments/list", departments.List)
	root.GET("/department/new", departments.New)
	root.POST("/department/create", departments.Create)
	root.GET("/edit/{department_id}", departments.Edit)
	root.GET("/show/{department_id}", departments.Show)
	root.PUT("/update/{department_id}", departments.Update)
	root.DELETE("/destroy/{department_id}", departments.Destroy)
	root.ServeFiles("/", http.FS(public.FS()))
}
