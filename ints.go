// Package ints provides some utilities around ints.
package ints

import (
	"fmt"
	"strconv"
)

// Parse is a simple wrapper around strconv.ParseInt.
//
// Parse considers "" to be 0.
func Parse(s string) (v int, err error) {
	if s == "" {
		// Special-case: consider empty string to be 0.
		v = 0
	} else {
		v64, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			return -1, err
		}
		v = int(v64)
	}
	return
}

// ParseBetween parses a string to int within an inclusive interval.
//
// ParseBetween considers "" to be 0.
func ParseBetween(s string, min, max int) (v int, err error) {
	v, err = Parse(s)
	if err != nil {
		return
	}
	if min > max {
		err = fmt.Errorf("min (%d) > max (%d)", min, max)
		return
	}
	if v < min {
		err = fmt.Errorf("wanted value >= %d, got %d", min, v)
		return
	}
	if v > max {
		err = fmt.Errorf("wanted value <= %d, got %d", max, v)
		return
	}
	return
}
