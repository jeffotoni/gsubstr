package gsubstr

import "testing"

func TestSubstr(t *testing.T) {
	type args struct {
		value string
		leni  int
		lenf  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test_substr_01", args{"DD-MM-YYYY", 0, 2}, "DD"},
		{"test_substr_02", args{"DD-MM-YYYY", 3, 2}, "MM"},
		{"test_substr_03", args{"DD-MM-YYYY", 6, 4}, "YYYY"},

		{"test_substr_04", args{"YYYYMMDD", 0, 4}, "YYYY"},
		{"test_substr_05", args{"YYYYMMDD", 4, 2}, "MM"},
		{"test_substr_06", args{"YYYYMMDD", 6, 2}, "DD"},

		{"test_substr_07", args{"ABCDEFGH", 0, 4}, "ABCD"},
		{"test_substr_08", args{"ABCDEFGH", 4, 1}, "E"},
		{"test_substr_08", args{"ABCDEFGH", 4, 0}, ""},
		{"test_substr_08", args{"ABCDEFGH", 4, 0}, ""},
		{"test_substr_08", args{"ABCDEFGH", 4, -1}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Substr(tt.args.value, tt.args.leni, tt.args.lenf); got != tt.want {
				t.Errorf("Substr() = %v, want %v", got, tt.want)
			}
		})
	}
}
