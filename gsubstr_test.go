package gsubstr

import (
	"fmt"
	"testing"
)

// This function is named ExampleSubstr()
// it with the Example .
func ExampleSubstr() {
	// This example Substr function accepts some types
	// specific func (value string, leni int, nums ...int).
	fmt.Println(Substr("DD-MM-YYYY", 0, 2))
	// Output -> DD

	fmt.Println(Substr("DD-MM-YYYY", 3, 2))
	// Output -> MM

	fmt.Println(Substr("ABCDEFGH", -4, 2))
	// Output -> EF

	fmt.Println(Substr("ABCD", -1, -5))
	// Output -> D

	fmt.Println(Substr("ABCDEFGH", -4, 3))
	// Output -> EFG

	fmt.Println(Substr("ABCDEFGH", 2, -3))
	// Output -> CDE

	fmt.Println(Substr("ABCDEFGH", 3, -3))
	// Output -> DE

	fmt.Println(Substr("ABCDEFGH", 2, -1))
	// Output -> CDEFG

	fmt.Println(Substr("ABCDEFGH", -6, -3))
	// Output -> CDE

	fmt.Println(Substr("ABCDEFGH", -6, -1))
	// Output -> CDEFG
}

func TestSubstrZero(t *testing.T) {
	s := Substr("", 0, 0, 0, 0)
	if s != "" {
		t.Errorf("Error Substr zero want len == 0 got:%s", s)
	}
}

func TestSubstrNegative(t *testing.T) {
	s := Substr("ABCD", -1, -5)
	if s != "D" {
		t.Errorf("Error Substr Negative want D got:%s", s)
	}
	t.Log(s)
}

func TestSubstr(t *testing.T) {
	type args struct {
		value string
		leni  int
		nums  []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.

		{"test_substr_", args{"DD-MM-YYYY", 0, []int{0}}, ""},
		{"test_substr_", args{"DD-MM-YYYY", 0, []int{2}}, "DD"},
		{"test_substr_", args{"DD-MM-YYYY", 3, []int{2}}, "MM"},
		{"test_substr_", args{"DD-MM-YYYY", 6, []int{4}}, "YYYY"},

		{"test_substr_", args{"YYYYMMDD", 0, []int{4}}, "YYYY"},
		{"test_substr_", args{"YYYYMMDD", 4, []int{2}}, "MM"},
		{"test_substr_", args{"YYYYMMDD", 6, []int{2}}, "DD"},

		{"test_substr_", args{"ABCDEFGH", 0, []int{8}}, "ABCDEFGH"},
		{"test_substr_", args{"ABCDEFGH", 8, []int{8}}, ""},

		{"test_substr_", args{"ABCDEFGH", 0, []int{4}}, "ABCD"},
		{"test_substr_", args{"ABCDEFGH", 4, []int{1}}, "E"},
		{"test_substr_", args{"ABCDEFGH", 4, []int{0}}, ""},
		{"test_substr_", args{"ABCDEFGH", 4, []int{2}}, "EF"},

		{"test_substr_", args{"ABCDEFGH", 0, []int{}}, "ABCDEFGH"},
		{"test_substr_", args{"ABCDEFGH", 1, []int{}}, "BCDEFGH"},
		{"test_substr_", args{"ABCDEFGH", 2, []int{}}, "CDEFGH"},
		{"test_substr_", args{"ABCDEFGH", 3, []int{}}, "DEFGH"},
		{"test_substr_", args{"ABCDEFGH", 4, []int{}}, "EFGH"},
		{"test_substr_", args{"ABCDEFGH", 5, []int{}}, "FGH"},
		{"test_substr_", args{"ABCDEFGH", 6, []int{}}, "GH"},
		{"test_substr_", args{"ABCDEFGH", 7, []int{}}, "H"},
		{"test_substr_", args{"ABCDEFGH", 8, []int{}}, ""},
		{"test_substr_", args{"ABCDEFGH", 9, []int{}}, ""},

		{"test_substr_", args{"ABCDEFGH", -1, []int{}}, "H"},
		{"test_substr_", args{"ABCDEFGH", -2, []int{}}, "GH"},
		{"test_substr_", args{"ABCDEFGH", -3, []int{}}, "FGH"},
		{"test_substr_", args{"ABCDEFGH", -4, []int{}}, "EFGH"},
		{"test_substr_", args{"ABCDEFGH", -5, []int{}}, "DEFGH"},
		{"test_substr_", args{"ABCDEFGH", -8, []int{}}, "ABCDEFGH"},
		{"test_substr_", args{"ABCDEFGH", -9, []int{}}, ""},

		{"test_substr_", args{"ABCDEFGH", -1, []int{2}}, "H"},
		{"test_substr_", args{"ABCDEFGH", -1, []int{3}}, "H"},
		{"test_substr_", args{"ABCDEFGH", -1, []int{8}}, "H"},
		{"test_substr_", args{"ABCDEFGH", -1, []int{10}}, "H"},

		{"test_substr_", args{"ABCDEFGH", -2, []int{1}}, "G"},
		{"test_substr_", args{"ABCDEFGH", -2, []int{2}}, "GH"},
		{"test_substr_", args{"ABCDEFGH", -3, []int{1}}, "F"},
		{"test_substr_", args{"ABCDEFGH", -3, []int{2}}, "FG"},
		{"test_substr_", args{"ABCDEFGH", -4, []int{1}}, "E"},
		{"test_substr_", args{"ABCDEFGH", -4, []int{2}}, "EF"},
		{"test_substr_", args{"ABCDEFGH", -4, []int{3}}, "EFG"},

		{"test_substr_", args{"ABCDEFGH", 0, []int{-1}}, "ABCDEFG"},
		{"test_substr_", args{"ABCDEFGH", 1, []int{-1}}, "BCDEFG"},
		{"test_substr_", args{"ABCDEFGH", 2, []int{-1}}, "CDEFG"},
		{"test_substr_", args{"ABCDEFGH", 3, []int{-1}}, "DEFG"},

		{"test_substr_", args{"ABCDEFGH", 0, []int{-0}}, ""},
		{"test_substr_", args{"ABCDEFGH", 1, []int{-0}}, ""},

		{"test_substr_", args{"ABCDEFGH", 1, []int{-2}}, "BCDEF"},
		{"test_substr_", args{"ABCDEFGH", 1, []int{-3}}, "BCDE"},
		{"test_substr_", args{"ABCDEFGH", 2, []int{-3}}, "CDE"},
		{"test_substr_", args{"ABCDEFGH", 3, []int{-3}}, "DE"},
		{"test_substr_", args{"ABCDEFGH", 1, []int{-3}}, "BCDE"},
		{"test_substr_", args{"ABCDEFGH", 2, []int{-3}}, "CDE"},
		{"test_substr_", args{"ABCDEFGH", 2, []int{-2}}, "CDEF"},
		{"test_substr_", args{"ABCDEFGH", 2, []int{-1}}, "CDEFG"},
		{"test_substr_", args{"ABCDEFGH", 2, []int{0}}, ""},

		{"test_substr_", args{"ABCDEFGH", -1, []int{1}}, "H"},
		{"test_substr_", args{"ABCDEFGH", -1, []int{5}}, "H"},

		{"test_substr_", args{"ABCDEFGH", -2, []int{2}}, "GH"},
		{"test_substr_", args{"ABCDEFGH", -2, []int{5}}, "GH"},

		{"test_substr_", args{"ABCDEFGH", -8, []int{0}}, ""},
		{"test_substr_", args{"ABCDEFGH", -8, []int{-0}}, ""},
		{"test_substr_", args{"ABCDEFGH", -8, []int{-1}}, "ABCDEFG"},
		{"test_substr_", args{"ABCDEFGH", -10, []int{-2}}, "ABCDEF"},
		{"test_substr_", args{"ABCDEFGH", -7, []int{-3}}, "BCDE"},
		{"test_substr_", args{"ABCDEFGH", -7, []int{-4}}, "BCD"},
		{"test_substr_", args{"ABCDEFGH", -7, []int{-5}}, "BC"},
		{"test_substr_", args{"ABCDEFGH", -7, []int{-6}}, "B"},
		{"test_substr_", args{"ABCDEFGH", -7, []int{-7}}, ""},

		{"test_substr_", args{"ABCDEFGH", -6, []int{-7}}, ""},
		{"test_substr_", args{"ABCDEFGH", -6, []int{-6}}, ""},
		{"test_substr_", args{"ABCDEFGH", -6, []int{-5}}, "C"},
		{"test_substr_", args{"ABCDEFGH", -6, []int{-4}}, "CD"},
		{"test_substr_", args{"ABCDEFGH", -6, []int{-3}}, "CDE"},
		{"test_substr_", args{"ABCDEFGH", -6, []int{-2}}, "CDEF"},
		{"test_substr_", args{"ABCDEFGH", -6, []int{-1}}, "CDEFG"},
		{"test_substr_", args{"ABCDEFGH", -6, []int{-0}}, ""},

		{"test_substr_", args{"A", 0, []int{1}}, "A"},
		{"test_substr_", args{"A", 0, []int{2}}, "A"},
		{"test_substr_", args{"A", 0, []int{-1}}, ""},
		{"test_substr_", args{"AB", 0, []int{-1}}, "A"},
		{"test_substr_", args{"ABC", 1, []int{-1}}, "B"},
		{"test_substr_", args{"ABCD", -1, []int{-5}}, "D"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Substr(tt.args.value, tt.args.leni, tt.args.nums...); got != tt.want {
				t.Errorf("Substr() = %v, want %v", got, tt.want)
			}
		})
	}
}
