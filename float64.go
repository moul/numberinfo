package numberinfo

import (
	"fmt"
	"math"
	"math/big"
)

// Float64Number implements the Number floaterface for base type float64
type Float64Number struct {
	value float64
	int64 *Int64Number
}

// Float64 returns an Float64Number instance
func Float64(value float64) *Float64Number {
	return &Float64Number{value: value}
}

// Int64 returns the equivalent Int64Number object
func (n *Float64Number) Int64() (*Int64Number, error) {
	if n.value >= math.MaxInt64 {
		return nil, fmt.Errorf("int64 overflow")
	}
	if n.value <= math.MinInt64 {
		return nil, fmt.Errorf("int64 underflow")
	}

	return Int64(int64(n.value)), nil
}

// BigFactorial returns the factorial value as a *big.Float
func (n *Float64Number) BigFactorial() (*big.Int, error) {
	int64, err := n.Int64()
	if err != nil {
		return nil, err
	}
	return int64.BigFactorial()
}
