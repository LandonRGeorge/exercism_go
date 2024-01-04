package cipher

import (
	"testing"
)

func TestExpandKey(t *testing.T) {
	key := "lemon"
	input := "attackatdawn"
	want := "lemonlemonle"
	got := expandKey(key, input)
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
