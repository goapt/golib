package number

import (
	"testing"
)

func TestUnique(t *testing.T) {
	for i := 0; i < 100; i++ {
		if len(Unique("")) != 20 {
			t.Error("number unique length error")
		}
	}
}
