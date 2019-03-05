package dividerapi

import (
	"context"
	"fmt"
	"log"

	divider "goa.design/goa/examples/error/gen/divider"
)

// divider service example implementation.
// The example methods log the requests and return zero values.
type dividersrvc struct {
	logger *log.Logger
}

// NewDivider returns the divider service implementation.
func NewDivider(logger *log.Logger) divider.Service {
	return &dividersrvc{logger}
}

// IntegerDivide implements integer_divide.
func (s *dividersrvc) IntegerDivide(ctx context.Context, p *divider.IntOperands) (int, error) {
	if p.B == 0 {
		return 0, divider.MakeDivByZero(fmt.Errorf("right operand cannot be 0"))
	}
	if p.A%p.B != 0 {
		return 0, divider.MakeHasRemainder(fmt.Errorf("remainder is %d", p.A%p.B))
	}
	return p.A / p.B, nil
}

// Divide implements divide.
func (s *dividersrvc) Divide(ctx context.Context, p *divider.FloatOperands) (float64, error) {
	if p.B == 0 {
		return 0, divider.MakeDivByZero(fmt.Errorf("right operand cannot be 0"))
	}
	return p.A / p.B, nil
}
