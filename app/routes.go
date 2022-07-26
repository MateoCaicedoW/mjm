package app

import (
	"net/http"

	"mjm/public"
	"mjm/app/actions/home"
	"mjm/app/middleware"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.RequestID)
	root.Use(middleware.Database)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/", home.Index)
	root.ServeFiles("/", http.FS(public.FS()))
}