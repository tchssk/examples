package calc

import (
	"context"
	"log"
	"os"
	"testing"

	calcsvc "goa.design/examples/basic/gen/calc"
)

func TestCalcAdd(t *testing.T) {
	logger := log.New(os.Stdout, "", log.Ltime|log.Lshortfile)
	calcSvc := NewCalc(logger)

	cases := map[string]struct {
		a        int
		b        int
		expected int
	}{
		"1+1":   {a: 1, b: 1, expected: 2},
		"1+10":  {a: 1, b: 10, expected: 11},
		"1+100": {a: 1, b: 100, expected: 101},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			payload := &calcsvc.AddPayload{
				A: tc.a,
				B: tc.b,
			}
			actual, err := calcSvc.Add(context.Background(), payload)
			if err != nil {
				t.Errorf("got error %v, expected none", err)
			}
			if actual != tc.expected {
				t.Errorf("got %v, expected %v", actual, tc.expected)
			}
		})
	}
}
