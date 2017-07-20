package main

import (
	"context"
	"net/http"

	"github.com/goadesign/examples/security/app"
	"github.com/goadesign/goa"
)

// NewBasicAuthMiddleware creates a middleware that checks for the presence of a basic auth header
// and validates its content.
func NewBasicAuthMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			authenticated := true
			// Retrieve and log basic auth info
			user, pass, ok := req.BasicAuth()
			// A real app would do something more interesting here
			if !ok {
				rd := goa.ContextRequest(ctx)
				if rd.URL.Path != "/basic" {
					goa.LogInfo(ctx, "failed basic auth")
					return ErrUnauthorized("missing auth")
				}
				authenticated = false
			}

			// Proceed
			goa.LogInfo(ctx, "basic", "user", user, "pass", pass)
			ctx = context.WithValue(ctx, authContextKey, authenticated)
			return h(ctx, rw, req)
		}
	}
}

type contextKey string

const authContextKey contextKey = "Authenticated"

// BasicController implements the BasicAuth resource.
type BasicController struct {
	*goa.Controller
}

// NewBasicController creates a BasicAuth controller.
func NewBasicController(service *goa.Service) *BasicController {
	return &BasicController{Controller: service.NewController("BasicController")}
}

// General runs the general action.
func (c *BasicController) General(ctx *app.GeneralBasicContext) error {
	authenticated, ok := ctx.Context.Value(authContextKey).(bool)
	if ok && authenticated {
		crtx, err := app.NewSecureBasicContext(ctx.Context, ctx.RequestData.Request, ctx.ResponseData.Service)
		if err != nil {
			return ctx.InternalServerError()
		}
		return c.Secure(crtx)
	}
	crtx, err := app.NewUnsecureBasicContext(ctx.Context, ctx.RequestData.Request, ctx.ResponseData.Service)
	if err != nil {
		return ctx.InternalServerError()
	}
	return c.Unsecure(crtx)
}

// Secure runs the secure action.
func (c *BasicController) Secure(ctx *app.SecureBasicContext) error {
	res := &app.Success{OK: true}
	ctx.ResponseData.Header().Set("Secured", "true")
	return ctx.OK(res)
}

// Unsecure runs the unsecure action.
func (c *BasicController) Unsecure(ctx *app.UnsecureBasicContext) error {
	res := &app.Success{OK: true}
	ctx.ResponseData.Header().Set("Secured", "false")
	return ctx.OK(res)
}
