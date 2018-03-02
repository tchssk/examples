// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "survey": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/goadesign/examples/form_encoding/design
// --out=$(GOPATH)/src/github.com/goadesign/examples/form_encoding
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// SubmitSurveyFormContext provides the survey_form submit action context.
type SubmitSurveyFormContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *YesNoPayload
}

// NewSubmitSurveyFormContext parses the incoming request URL and body, performs validations and creates the
// context used by the survey_form controller submit action.
func NewSubmitSurveyFormContext(ctx context.Context, r *http.Request, service *goa.Service) (*SubmitSurveyFormContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := SubmitSurveyFormContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *SubmitSurveyFormContext) OK(r *ResultMedia) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.example.form")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
