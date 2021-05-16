// Code generated by goa v3.4.0, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package client

import (
	"encoding/json"
	"fmt"

	calc "goa.design/examples/error/gen/calc"
)

// BuildDividePayload builds the payload for the calc divide endpoint from CLI
// flags.
func BuildDividePayload(calcDivideBody string) (*calc.DividePayload, error) {
	var err error
	var body DivideRequestBody
	{
		err = json.Unmarshal([]byte(calcDivideBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"dividend\": 6322633713974661021,\n      \"divisor\": 3793862871819669726\n   }'")
		}
	}
	v := &calc.DividePayload{
		Dividend: body.Dividend,
		Divisor:  body.Divisor,
	}

	return v, nil
}