package number

import (
	"testing"
)

func TestNewRing(t *testing.T) {
	ring := NewRing(10, 100)
	for i := 0; i <= 500; i++ {
		if ring.Next() < 10 || ring.Next() > 100 {
			t.Error("next ring number out of bounds")
		}
	}
}
