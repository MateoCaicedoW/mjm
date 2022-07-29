package app

import (
	"net/http"

	"mjm/app/actions"
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

	root.GET("/requirements", actions.ListRequirements)
	root.GET("/requirements/new", actions.RequirementsNew)
	root.ServeFiles("/", http.FS(public.FS()))
}
