package app

import (
	"net/http"

	"mjm/app/actions"
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
	root.GET("/requirement-type", actions.List)
	root.GET("/requirement-type-new", actions.New)
	root.POST("/requirement-type-create", actions.Create)
	

	root.ServeFiles("/", http.FS(public.FS()))
}
