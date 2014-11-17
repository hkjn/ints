// Package ints provides some utilities for ints.
package ints // import "hkjn.me/ints"

import (
	"fmt"
	"strconv"
)

// ParseStrings parses multiple strings as int.
func ParseStrings(in ...string) ([]int, error) {
	result := make([]int, len(in))
	for i, s := range in {
		r, err := Parse(s)
		if err != nil {
			return []int{}, err
		}
		result[i] = r
	}
	return result, nil
}

// Parse is a simple wrapper around strconv.ParseInt.
func Parse(s string) (int, error) {
	v64, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return -1, err
	}
	return int(v64), nil
}

// ParseWithDefault is like Parse but considers "" to be default value.
func ParseWithDefault(s string, def int) (int, error) {
	if s == "" {
		return def, nil
	} else {
		return Parse(s)
	}
}

// ParseBetween parses a string to int within an inclusive interval.
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

// ParseBetweenWithDefault is like ParseBetween but considers "" to be 0.
func ParseBetweenWithDefault(s string, min, max int) (v int, err error) {
	if s == "" {
		s = "0"
	}
	return ParseBetween(s, min, max)
}
