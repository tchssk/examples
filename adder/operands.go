//go:generate goagen bootstrap -d github.com/goadesign/examples/adder/design
package main

import (
	"strconv"

	"github.com/goadesign/examples/adder/app"
	"github.com/goadesign/goa"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("operands")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	if ctx.Left == 0 && ctx.Right == 0 {
		ctx.TooManyRequests()
	}
	return ctx.OK([]byte(strconv.Itoa(ctx.Left + ctx.Right)))
}
