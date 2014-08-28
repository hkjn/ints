// ints_test.go; Tests for package ints.
package ints

import (
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{"0", 0},
		{"-0", 0},
		{"-2", -2},
		{"4294967296", 4294967296},
		{"4294967296", 1 << 32},
		{"4611686018427387904", 1 << 62},
	}
	for i, tt := range cases {
		got, err := Parse(tt.s)
		if err != nil {
			t.Errorf("[%d] Parse(%q) got error %q, want %v\n", i, tt.s, err, tt.want)
		} else if got != tt.want {
			t.Errorf("[%d] Parse(%q) => %v, want %v\n", i, tt.s, got, tt.want)
		}
		// TODO: test failure cases
	}
}

func TestParseWithDefault(t *testing.T) {
	cases := []struct {
		s    string
		def  int
		want int
	}{
		{"", -1, -1},
		{"", 1, 1},
		{"", 0, 0},
		{"0", -1, 0},
		{"0", 1, 0},
		{"4294967296", -1, 1 << 32},
		{"4611686018427387904", 1 << 62, 1 << 62},
	}
	for i, tt := range cases {
		got, err := ParseWithDefault(tt.s, tt.def)
		if err != nil {
			t.Errorf("[%d] ParseWithDefault(%q, %d) got error %q, want %v\n", i, tt.s, tt.def, err, tt.want)
		} else if got != tt.want {
			t.Errorf("[%d] ParseWithDefault(%q, %d) => %v, want %v\n", i, tt.s, tt.def, got, tt.want)
		}
		// TODO: test failure cases
	}
}

func TestParseBetween(t *testing.T) {
	cases := []struct {
		s        string
		min, max int
		want     int
	}{
		{"0", 0, 0, 0},
		{"-1", -100, 100, -1},
		{"129", -100, 129, 129},
	}
	for i, tt := range cases {
		got, err := ParseBetween(tt.s, tt.min, tt.max)
		if err != nil {
			t.Errorf("[%d] ParseBetween(%q, %d, %d) got error %q, want %v\n", i, tt.s, tt.min, tt.max, err, tt.want)
		} else if got != tt.want {
			t.Errorf("[%d] ParseBetween(%q, %d, %d) => %v, want %v\n", i, tt.s, tt.min, tt.max, got, tt.want)
		}
		// TODO: test failure cases
	}
}
