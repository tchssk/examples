// Code generated by goa v3.4.3, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package client

import (
	"fmt"
	"strconv"

	calc "goa.design/examples/error/gen/calc"
)

// BuildDivPayload builds the payload for the calc div endpoint from CLI flags.
func BuildDivPayload(calcDivA string, calcDivB string) (*calc.DivPayload, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(calcDivA, 10, 64)
		a = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(calcDivB, 10, 64)
		b = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for b, must be INT")
		}
	}
	v := &calc.DivPayload{}
	v.A = a
	v.B = b

	return v, nil
}
