package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type m struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	Status bool   `db:"status"`
}

func TestStructToMapInterface(t *testing.T) {
	tm := &m{
		Id:     1,
		Name:   "test",
		Status: true,
	}
	m2 := StructToMapInterface(tm)
	m3 := map[string]interface{}{
		"id":     1,
		"name":   "test",
		"status": true,
	}
	assert.Equal(t, m2, m3)
}

func TestStructToMapString(t *testing.T) {
	tm := &m{
		Id:     1,
		Name:   "test",
		Status: true,
	}
	m2 := StructToMapString(tm)
	m3 := map[string]string{
		"id":     "1",
		"name":   "test",
		"status": "true",
	}
	assert.Equal(t, m2, m3)
}
