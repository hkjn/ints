// Tests for ints.
package ints

import (
	"fmt"
	"testing"
)

type ParseTest struct {
	In  string
	Out int
}

func TestParse(t *testing.T) {
	cases := []ParseTest{
		ParseTest{"", 0},
		ParseTest{"0", 0},
		ParseTest{"-0", 0},
		ParseTest{"-2", -2},
		ParseTest{"4294967296", 4294967296},
		ParseTest{"4294967296", 1 << 32},
		ParseTest{"4611686018427387904", 1 << 62},
	}
	for i, c := range cases {
		actual, _ := Parse(c.In)
		if actual != c.Out {
			t.Fatalf("[%d] Parse(%v) want %v, got %v\n", i, c.In, c.Out, actual)
		}
		// TODO: test failure cases
	}
}

type ParseBetweenIn struct {
	s        string
	min, max int
}

func (p ParseBetweenIn) String() string {
	return fmt.Sprintf("%q, %d, %d", p.s, p.min, p.max)
}

func TestParseBetween(t *testing.T) {
	cases := map[ParseBetweenIn]int{
		ParseBetweenIn{"", 0, 0}:         0,
		ParseBetweenIn{"0", 0, 0}:        0,
		ParseBetweenIn{"-1", -100, 100}:  -1,
		ParseBetweenIn{"129", -100, 129}: 129,
	}
	for k, out := range cases {
		actual, err := ParseBetween(k.s, k.min, k.max)
		if err != nil {
			t.Fatalf("ParseBetween(%v) got error: %v\n", k, err)
		}
		if actual != out {
			t.Fatalf("ParseBetween(%v) want %v, got %v\n", k, out, actual)
		}
		// TODO: test failure cases
	}
}
