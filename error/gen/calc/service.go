// Code generated by goa v3.4.3, DO NOT EDIT.
//
// calc service
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Service is the calc service interface.
type Service interface {
	// Div implements div.
	Div(context.Context, *DivPayload) (res int, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "calc"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"div"}

// DivPayload is the payload type of the calc service div method.
type DivPayload struct {
	// Left operand
	A int
	// Right operand
	B int
}

// MakeDivByZero builds a goa.ServiceError from an error.
func MakeDivByZero(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "DivByZero",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
