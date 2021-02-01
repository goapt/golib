package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrMap_Keys(t *testing.T) {
	m := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
		"d": "4",
	}

	assert.Equalf(t, len(StrMap(m).Keys()), 4, "StrMap Keys return not abcd")
}

func TestStrMap_ToInterface(t *testing.T) {
	m := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
		"d": "4",
	}

	m2 := StrMap(m).ToInterface()

	assert.Equal(t, m2["a"], "1", "StrMap to interface error")
}

func TestStrMaps_IndexBy(t *testing.T) {

	m := make([]map[string]string, 0)
	m = append(m, map[string]string{
		"id":   "1",
		"name": "test1",
	})
	m = append(m, map[string]string{
		"id":   "2",
		"name": "test2",
	})
	m = append(m, map[string]string{
		"id":   "3",
		"name": "test3",
	})
	m = append(m, map[string]string{
		"id":   "4",
		"name": "test4",
	})

	m2 := StrMaps(m).IndexBy("id")
	m3 := StrMaps(m).IndexBy("id", "id")

	assert.Equal(t, m2["1"]["name"], "test1", "StrMaps indexBy key error", m2)
	assert.Equal(t, m3["1"]["id"], "1", "StrMaps indexBy key error", m3)

}

func TestNewMapSorter(t *testing.T) {
	m := map[string]string{
		"c":  "c",
		"a":  "a",
		"1":  "1",
		"b":  "b",
		"-1": "-1",
	}

	sort := NewMapSorter(m)
	sort.Sort()

	m2 := map[string]string{
		"-1": "-1",
		"1":  "1",
		"a":  "a",
		"b":  "b",
		"c":  "c",
	}

	assert.Equal(t, m, m2)
}
