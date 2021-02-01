package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiffSlice(t *testing.T) {
	a := []string{"a", "b", "c", "d"}
	b := []string{"b", "d"}

	diff := DiffSlice(a, b)

	assert.Equal(t, diff, []string{"a", "c"}, "DiffSlice error")
}

func TestMinSlice(t *testing.T) {
	tests := []struct {
		name  string
		args  []int64
		wantM int64
	}{
		{
			"min",
			[]int64{3, 2, 6, 4, 7, 8, 9, 4},
			2,
		},
		{
			"min signed",
			[]int64{3, 2, 6, -1, 4, 7, 8, 9, 4},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := MinSlice(tt.args); gotM != tt.wantM {
				t.Errorf("MinSlice() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func TestMaxSlice(t *testing.T) {
	tests := []struct {
		name  string
		args  []int64
		wantM int64
	}{
		{
			"max",
			[]int64{3, 2, 6, 4, 7, 8, 9, 4},
			9,
		},
		{
			"max signed",
			[]int64{3, 2, 6, -1, 4, 7, 8, 9, 4},
			9,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := MaxSlice(tt.args); gotM != tt.wantM {
				t.Errorf("MaxSlice() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}
