package app

import (
	"net/http"

	"mjm/app/actions/requirements"
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

	root.GET("/requirements", requirements.List)
	root.GET("/requirements/new", requirements.New)
	root.POST("/requirements/new", requirements.Create)
	root.ServeFiles("/", http.FS(public.FS()))
}
