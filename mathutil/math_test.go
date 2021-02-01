package mathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsInt(t *testing.T) {
	list := []int{-3, -2, -1, 0, 1, 2, 3}
	list2 := []int{3, 2, 1, 0, 1, 2, 3}

	for k, v := range list {
		if AbsInt(v) != list2[k] {
			t.Errorf("%d Abs is not equal to %d", v, list2[k])
		}
	}
}

func TestMaxInt(t *testing.T) {
	min := 1
	max := 2

	if MaxInt(min, max) != max {
		t.Errorf("not equal to max value")
	}
}

func TestMinInt(t *testing.T) {
	min := -1
	max := 2

	if MinInt(min, max) != min {
		t.Errorf("not equal to min value")
	}
}

func TestMaxInt64(t *testing.T) {
	min := int64(1)
	max := int64(2)

	if MaxInt64(min, max) != max {
		t.Errorf("not equal to max value")
	}
}

func TestMinInt64(t *testing.T) {
	min := int64(-1)
	max := int64(2)

	if MinInt64(min, max) != min {
		t.Errorf("not equal to min value")
	}
}

func TestPowInt(t *testing.T) {
	a := PowInt(2, 0)
	assert.Equal(t, 1, a)
	b := PowInt(10, 2)
	assert.Equal(t, 100, b)
	c := PowInt(2, -1)
	assert.Equal(t, 1, c)
}
