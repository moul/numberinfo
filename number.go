package numberinfo

type Number struct {
	int64 int64
}

func Int64(input int64) Number {
	return Number{
		int64: input,
	}
}

func (n *Number) Int64() int64 {
	return n.int64
}
