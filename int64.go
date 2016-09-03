package numberinfo

import "math/big"

// Int64Number implements the Number interface for base type int64
type Int64Number struct {
	value int64
}

// Int64 returns an Int64Number instance
func Int64(value int64) *Int64Number {
	return &Int64Number{value: value}
}

// Float64Number returns the equivalent Int64Number object
func (n *Int64Number) Float64Number() (*Float64Number, error) {
	return Float64(float64(n.value)), nil
}

// Int64Number returns itself
func (n *Int64Number) Int64Number() (*Int64Number, error) {
	return n, nil
}

// BigFactorial returns the factorial value as a *big.Int
func (n *Int64Number) BigFactorial() (*big.Int, error) {
	x := new(big.Int)
	x.MulRange(1, n.value)
	return x, nil
}

// IsPrime returns true if the number is prime
func (n *Int64Number) IsPrime() bool {
	if n.value < 2 {
		return false
	}

	for i := int64(2); i < n.value; i++ {
		if n.value%i == 0 {
			return false
		}
	}
	return true
}
