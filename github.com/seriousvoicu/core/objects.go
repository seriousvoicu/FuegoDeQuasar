package core

import ()

type Vector2 struct {
	X, Y  float64
	empty bool
}

func (element Vector2) IsEmpty() bool {
	return element.empty
}

func GetEmptyVector() *Vector2 {
	return &Vector2{0, 0, true}
}
