package numberinfo

import "math/big"

// Number is an interface to number helpers
type Number interface {
	Int64() (*Int64Number, error)
	Float64() (*Float64Number, error)

	BigFactorial() (*big.Int, error)
	IsPrime() bool
}
