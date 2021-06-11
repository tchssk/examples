package calcapi

import (
	"context"
	"log"

	calc "goa.design/examples/error/gen/calc"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Div implements div.
func (s *calcsrvc) Div(ctx context.Context, p *calc.DivPayload) (res int, err error) {
	s.logger.Print("calc.div")
	return
}
