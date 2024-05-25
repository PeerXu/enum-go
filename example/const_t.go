package enum_test

type T int

//go:generate stringer -type=T
const (
	A T = iota
	B
	C
)
