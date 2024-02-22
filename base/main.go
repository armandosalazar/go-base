package base

import "math"

func Hi(name string) string {
	return "Hi {name}"
}

func Vars() {
	// @note ":=" is used to declare and initialize a variable
	a := true
	// @note "var" is used to declare a variable
	var b bool

	maxIntergerUint8 := math.MaxUint8
	println(maxIntergerUint8)

	if a && b {
		println("a && b")
	} else {
		println("!(a && b)")
	}
}
