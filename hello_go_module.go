package hello_go_module

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes in two [Number] parameters and returns the sum of the parameters.
//
// Try it out: [mathisfun/addition]
//
// [mathisfun/addition]: https://www.mathsisfun.com/numbers/addition.html
func Add[N Number](a, b N) N {
	return a + b
}
