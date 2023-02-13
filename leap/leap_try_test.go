package leap_test

import (
	"leap"
	"testing"
)

func TestIsLeapYear(t *testing.T) {
	testCases := []struct {
		arg  int
		want bool
	}{
		{arg: 1996, want: true},
		{arg: 1900, want: false},
		{arg: 1600, want: true},
		{arg: 2000, want: true},
		{arg: 2100, want: false},
	}
	for _, tc := range testCases {
		got := leap.IsLeapYear(tc.arg)
		if got != tc.want {
			t.Errorf("IsLeapYear(%d), got %t want %t", tc.arg, got, tc.want)
		}
	}
}
