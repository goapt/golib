package filesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteSize(t *testing.T) {
	assert.EqualValues(t, MustParse("1"), 1)
	assert.EqualValues(t, MustParse("1b"), 1)
	assert.EqualValues(t, MustParse("1k"), KB)
	assert.EqualValues(t, MustParse("1m"), MB)
	assert.EqualValues(t, MustParse("1g"), GB)
	assert.EqualValues(t, MustParse("1t"), TB)
	assert.EqualValues(t, MustParse("1p"), PB)

	assert.EqualValues(t, MustParse(" -1"), -1)
	assert.EqualValues(t, MustParse(" -1 b"), -1)
	assert.EqualValues(t, MustParse(" -1 kb "), -1*KB)
	assert.EqualValues(t, MustParse(" -1 mb "), -1*MB)
	assert.EqualValues(t, MustParse(" -1 gb "), -1*GB)
	assert.EqualValues(t, MustParse(" -1 tb "), -1*TB)
	assert.EqualValues(t, MustParse(" -1 pb "), -1*PB)

	assert.EqualValues(t, MustParse(" 1.5"), 1)
	assert.EqualValues(t, MustParse(" 1.5 kb "), 1.5*KB)
	assert.EqualValues(t, MustParse(" 1.5 mb "), 1.5*MB)
	assert.EqualValues(t, MustParse(" 1.5 gb "), 1.5*GB)
	assert.EqualValues(t, MustParse(" 1.5 tb "), 1.5*TB)
	assert.EqualValues(t, MustParse(" 1.5 pb "), 1.5*PB)
}

func TestByteSizeError(t *testing.T) {
	var err error
	_, err = Parse("--1")
	assert.Equal(t, err, ErrBadByteSize)
	_, err = Parse("hello world")
	assert.Equal(t, err, ErrBadByteSize)
	_, err = Parse("123.132.32")
	assert.Equal(t, err, ErrBadByteSize)
	_, err = Parse("120000000000000000000000000tb")
	assert.Equal(t, err, ErrBadByteSize)
}

func TestInt64_AsInt(t *testing.T) {
	b := Int64.AsInt(1)
	assert.Equal(t, 1, b)
}

func TestInt64_Int64(t *testing.T) {
	b := Int64(1)
	assert.Equal(t, int(1), b.AsInt())
}

func TestInt64_HumanString(t *testing.T) {
	tests := []struct {
		name string
		args int64
		want string
	}{
		{
			name: "byte",
			args: 1,
			want: "1",
		},
		{
			name: "byteToKB",
			args: 1 * KB,
			want: "1024",
		},
		{
			name: "KBToMB",
			args: 1 * MB,
			want: "1024.00kb",
		},
		{
			name: "MBToGB",
			args: 1 * GB,
			want: "1024.00mb",
		},
		{
			name: "GBToTB",
			args: 1 * TB,
			want: "1024.00gb",
		},
		{
			name: "TBToPB",
			args: 1 * PB,
			want: "1024.00tb",
		},
		{
			name: "PB",
			args: 1 * 1024 * PB,
			want: "1024.00pb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Int64(tt.args)
			assert.Equal(t, tt.want, b.HumanString())
		})
	}

}

func TestInt64_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		args    int64
		want    string
		wantErr bool
	}{
		{name: "zero",
			args:    0,
			want:    "0",
			wantErr: false,
		},
		{name: "kb",
			args:    1 * KB,
			want:    "1kb",
			wantErr: false,
		},
		{name: "mb",
			args:    1 * MB,
			want:    "1mb",
			wantErr: false,
		},
		{name: "gb",
			args:    1 * GB,
			want:    "1gb",
			wantErr: false,
		},
		{name: "tb",
			args:    1 * TB,
			want:    "1tb",
			wantErr: false,
		},
		{name: "pb",
			args:    1 * PB,
			want:    "1pb",
			wantErr: false,
		},
		{name: "byte to <0",
			args:    -1000,
			want:    "-1000",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Int64(tt.args)
			text, err := b.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("%s in error", tt.name)
			}
			assert.Equal(t, tt.want, string(text))
		})
	}
}

func TestInt64_UnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		args    int64
		want    string
		wantErr bool
	}{
		{name: "success",
			args:    -1000,
			want:    "-1000",
			wantErr: false,
		},
		{name: "error",
			args:    100,
			want:    "127.0.0.1",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Int64(tt.args)
			err := b.UnmarshalText([]byte(tt.want))
			if (err != nil) != tt.wantErr {
				t.Errorf("%s in error", tt.name)
				return
			}
			if tt.wantErr {
				assert.Equal(t, err, ErrBadByteSize)
			} else {
				assert.Equal(t, tt.want, b.HumanString())
			}

		})
	}
}

func TestInt64_MustParse(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    int64
		wantErr bool
	}{
		{name: "success",
			args:    "1000",
			want:    1000,
			wantErr: false,
		},
		{name: "error",
			args:    "127.0.0.1",
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				assert.Panics(t, func() {
					b := MustParse(tt.args)
					assert.Equal(t, tt.want, b)
				})
			} else {
				b := MustParse(tt.args)
				assert.Equal(t, tt.want, b)
			}
		})
	}
}
