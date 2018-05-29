package mathutil

import (
	"testing"
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

func TestRound(t *testing.T) {
	list := []float64{-3.4, -0.8, 1, 1.3, 1.5, 1.6}
	list2 := []int{-3, -1, 1, 1, 2, 2}

	for k, v := range list {
		if vv, err := Round(v); vv != list2[k] || err != nil {
			t.Errorf("%f Round is not equal to %d, error: %s", v, list2[k], err)
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
