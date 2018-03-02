//go:generate goagen bootstrap -d github.com/goadesign/examples/form_encoding/design

package main

import (
	"github.com/goadesign/examples/form_encoding/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("survey")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "survey_form" controller
	c := NewSurveyFormController(service)
	app.MountSurveyFormController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
