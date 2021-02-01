package stringutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkGetRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString(8)
	}
}

func TestGetRandomString(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := RandomString(8)
		if len(s) != 8 {
			t.Error("string length error:" + s)
		}
	}
}

func TestTitleCasedName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{"userID"},
			want: "Userid",
		}, {
			name: "test2",
			args: args{"UserName"},
			want: "Username",
		}, {
			name: "test3",
			args: args{"user_id"},
			want: "UserId",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TitleCasedName(tt.args.name); got != tt.want {
				t.Errorf("TitleCasedName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnakeCasedName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{"UserId"},
			want: "user_id",
		}, {
			name: "test2",
			args: args{"userID"},
			want: "user_i_d",
		}, {
			name: "test3",
			args: args{"user_id"},
			want: "user_id",
		}, {
			name: "test4",
			args: args{"user_ID"},
			want: "user_i_d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeCasedName(tt.args.name); got != tt.want {
				t.Errorf("SnakeCasedName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustJsonEncode(t *testing.T) {
	m := map[string]interface{}{
		"id":     1,
		"name":   "test",
		"status": true,
	}
	s := MustJsonEncode(m)
	assert.Equal(t, `{"id":1,"name":"test","status":true}`, s)
}

func TestLeftpad(t *testing.T) {
	s := Leftpad("test", 10)
	assert.Equal(t, "      test", s)
}

func TestTrimBom(t *testing.T) {
	s := TrimBom(string([]byte{239, 187, 191}) + "abcd")
	assert.Equal(t, "abcd", s)
}
