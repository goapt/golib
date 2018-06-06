package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonTo_MapString(t *testing.T) {
	json := `{"a":1,"b":true,"c":"hello"}`

	m := JsonTo(json).MapString()

	assert.Equal(t,m["a"],"1","JsonTo MapString key a must value string 1:%v",m)
	assert.Equal(t,m["b"],"true","JsonTo MapString key b must value string true:%v",m)
	assert.Equal(t,m["c"],"hello","JsonTo MapString key c must value string hello:%v",m)
}