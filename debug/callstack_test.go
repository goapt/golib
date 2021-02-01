package debug

import (
	"testing"
)

func TestStack(t *testing.T) {
	b := Stack(2)

	if len(b) == 0 {
		t.Error("stack must > 0")
	}

	b2 := Stack(-1)

	if len(b2) == 0 {
		t.Error("stack must > 0")
	}
}
