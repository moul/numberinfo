package numberinfo

import "math/big"

// Number is an interface to number helpers
type Number interface {
	BigFactorial() (*big.Int, error)
}
