package collection

import "testing"

func TestInStr(t *testing.T) {
	m := []string{"a", "b", "c"}

	if !InStr("a", m) {
		t.Error("Item a not found!")
	}

	if InStr("d", m) {
		t.Error("Item d should not be found")
	}
}

func TestInInt(t *testing.T) {
	m := []int{1, 2, 3}

	if !InInt(1, m) {
		t.Error("Item 1 not found!")
	}

	if InInt(0, m) {
		t.Error("Item 0 should not be found")
	}
}

func TestInSlice(t *testing.T) {

	m := []int{1, 2, 3}

	if !InSlice(1, m) {
		t.Error("Item 1 not found!")
	}

	if InSlice(0, m) {
		t.Error("Item 0 should not be found")
	}

	m2 := []int64{1, 2, 3}

	if !InSlice(int64(1), m2) {
		t.Error("Item 1 not found!")
	}

	if InSlice(int64(0), m2) {
		t.Error("Item 0 should not be found")
	}

	m3 := []string{"1", "2", "3"}

	if !InSlice("1", m3) {
		t.Error("Item 1 not found!")
	}

	if InSlice("0", m3) {
		t.Error("Item 0 should not be found")
	}
}
