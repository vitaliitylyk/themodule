package themodule

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add returns the sum of two integers.
func Add[T Number](x, y T) T {
	return x + y
}
