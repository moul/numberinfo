package numberinfo

import "strconv"

type Number struct {
	int64   int64
	float64 float64
}

func Int64(input int64) *Number {
	return &Number{
		int64:   input,
		float64: float64(input),
	}
}

func Float64(input float64) *Number {
	return &Number{
		int64:   int64(input),
		float64: input,
	}
}

func (n *Number) Int64() int64 {
	return n.int64
}

func (n *Number) Int() int {
	return int(n.int64)
}

func (n *Number) Float64() float64 {
	return n.float64
}

func (n *Number) String() string {
	return strconv.Itoa(n.Int())
}
