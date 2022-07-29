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
	root.GET("/departments", departments.ListDepartment)
	root.GET("/add-department", departments.FormCreateDeparment)
	root.POST("/create-department", departments.CreateDepartment)
	root.GET("/edit/{department_id}", departments.Edit)
	root.GET("/view/{department_id}", departments.ViewDetails)
	root.PUT("/edit-data/{department_id}", departments.Editdepartment)
	root.DELETE("/delete/{department_id}", departments.DeleteDepartment)
	root.ServeFiles("/", http.FS(public.FS()))
}
