package models

import (
	"strconv"
)

const (
	W byte = 'w'
	X byte = 'x'
	Y byte = 'y'
	Z byte = 'z'
)

type Literal struct {
	Symbol   byte
	Validity bool
}

func ToString(literal Literal) string {
	return string(literal.Symbol) + " : " + strconv.FormatBool(literal.Validity)
}
