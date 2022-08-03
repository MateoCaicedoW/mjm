package app

import (
	"net/http"

	"mjm/app/actions/home"
	requirement_type "mjm/app/actions/requeriment_type"

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

	root.ServeFiles("/", http.FS(public.FS()))
}
