package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goapt/golib/timeutil"
)

func TestStrTo_Time(t *testing.T) {
	t1 := "2012-10-24 07:49:00"
	ct1 := StrTo(t1).MustTime().Format("2006-01-02 15:04:05")
	if t1 != ct1 {
		t.Error("StrTo_Time fail")
	}

	t2 := "2012-10-24T07:49:00+08:00"

	ct2, err := StrTo(t2).Time()

	if err != nil {
		t.Error(err)
	}

	ct3 := ct2.Format("2006-01-02 15:04:05")
	if ct3 != t1 {
		t.Error("StrTo_Time RFC3339 fail")
	}
}

func TestStrTo_Int(t *testing.T) {
	var i = "123"

	i8, err := StrTo(i).Int8()
	assert.NoError(t, err)
	assert.Equal(t, i8, int8(123))
	i16, err := StrTo(i).Int16()
	assert.NoError(t, err)
	assert.Equal(t, i16, int16(123))
	i32, err := StrTo(i).Int32()
	assert.NoError(t, err)
	assert.Equal(t, i32, int32(123))

	ui16, err := StrTo(i).Uint16()
	assert.NoError(t, err)
	assert.Equal(t, ui16, uint16(123))
	ui32, err := StrTo(i).Uint32()
	assert.NoError(t, err)
	assert.Equal(t, ui32, uint32(123))

	ii := StrTo(i).MustInt()
	assert.Equal(t, ii, 123)
	i64 := StrTo(i).MustInt64()
	assert.Equal(t, i64, int64(123))
	ui := StrTo(i).MustUint()
	assert.Equal(t, ui, uint(123))
	ui8 := StrTo(i).MustUint8()
	assert.Equal(t, ui8, uint8(123))
	ui64 := StrTo(i).MustUint64()
	assert.Equal(t, ui64, uint64(123))
}

func TestToStr(t *testing.T) {
	assert.Equal(t, "true", ToStr(true))
	assert.Equal(t, "1.2", ToStr(float32(1.2)))
	assert.Equal(t, "1.2", ToStr(float64(1.2)))
	assert.Equal(t, "123", ToStr(123))
	assert.Equal(t, "123", ToStr(int8(123)))
	assert.Equal(t, "123", ToStr(int16(123)))
	assert.Equal(t, "123", ToStr(int32(123)))
	assert.Equal(t, "123", ToStr(int64(123)))
	assert.Equal(t, "123", ToStr(uint(123)))
	assert.Equal(t, "123", ToStr(uint8(123)))
	assert.Equal(t, "123", ToStr(uint16(123)))
	assert.Equal(t, "123", ToStr(uint32(123)))
	assert.Equal(t, "123", ToStr(uint64(123)))
	assert.Equal(t, "123", ToStr("123"))
	assert.Equal(t, "123", ToStr([]byte("123")))

	assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", ToStr(timeutil.Zero()))
	var a interface{}
	a = 123
	assert.Equal(t, "123", ToStr(a))
}
