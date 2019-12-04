// Code generated by goa v2.0.8, DO NOT EDIT.
//
// calc gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/examples/basic/design -o
// $(GOPATH)/src/goa.design/examples/basic

package client

import (
	"encoding/json"
	"fmt"

	calc "goa.design/examples/basic/gen/calc"
	calcpb "goa.design/examples/basic/gen/grpc/calc/pb"
)

// BuildAddPayload builds the payload for the calc add endpoint from CLI flags.
func BuildAddPayload(calcAddMessage string) (*calc.AddPayload, error) {
	var err error
	var message calcpb.AddRequest
	{
		if calcAddMessage != "" {
			err = json.Unmarshal([]byte(calcAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"a\": 8133055152903002499,\n      \"b\": 3219793201326175278\n   }'")
			}
		}
	}
	v := &calc.AddPayload{
		A: int(message.A),
		B: int(message.B),
	}
	return v, nil
}
