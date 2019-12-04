package calc

import (
	"context"
	"log"

	calcsvc "goa.design/examples/basic/gen/calc"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcSvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calcsvc.Service {
	return &calcSvc{logger}
}

// Add implements add.
func (s *calcSvc) Add(ctx context.Context, p *calcsvc.AddPayload) (int, error) {
	if p.A == 1 {
		return 0, calcsvc.Accepted(1)
	}
	return p.A + p.B, nil
}
