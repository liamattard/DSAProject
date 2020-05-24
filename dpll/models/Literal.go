package models

import ()

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
