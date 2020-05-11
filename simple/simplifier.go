// Copyright (c) 2020, Peter Ohler, All rights reserved.

package simple

// Simplifier interface is for objects that can decompose themselves into
// simple data.
type Simplifier interface {

	// Simplify should return one of the simple types which are: nil, bool,
	// int64, float64, string, time.Time, []interface{}, or
	// map[string]interface{}.
	Simplify() interface{}
}