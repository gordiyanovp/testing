package task

import (
	"testing"
)

func TestSum(t *testing.T) {
	s := Sum(3, 6)
	if s != 9 {
		t.Errorf("exp %d, got %d", 9, s)
	}
}
