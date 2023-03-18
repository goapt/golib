package filesystem

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestIsExists(t *testing.T) {
	if !IsExists(".") {
		t.Error(". must exist")
		return
	}
}

func TestIsDir(t *testing.T) {
	if !IsDir(".") {
		t.Error(". should be a directory")
		return
	}
}

func TestReadLine(t *testing.T) {
	lchanError, err := ReadLine("../testdata/data1.txt")
	if err == nil || lchanError != nil {
		t.Fatal("ReadLine must be can not open file")
	}
	lchan, err := ReadLine("../testdata/data.txt")
	if err != nil {
		t.Fatal(err)
	}

	for line := range lchan {
		fmt.Println(string(line))
	}
}

func TestFile_IsWritable(t *testing.T) {
	var res bool
	res = IsWritable("../testdata/data.txt")
	assert.Equal(t, true, res)
	res = IsWritable("../testdata/data1.txt")
	assert.Equal(t, false, res)
}

func TestFile_IsReadable(t *testing.T) {
	res := IsReadable("../testdata/data.txt")
	assert.Equal(t, true, res)
	res = IsReadable("../testdata/data1.txt")
	assert.Equal(t, false, res)
}

func TestFile_IsExecutable(t *testing.T) {
	res := IsExecutable("../testdata/data.txt")
	assert.Equal(t, false, res)
}

func TestFile_IsDir(t *testing.T) {
	res := IsDir("../testdata")
	assert.Equal(t, true, res)
	res = IsDir("../testdata/data.txt")
	assert.Equal(t, false, res)
	res = IsDir("../testdata1")
	assert.Equal(t, false, res)
}

func TestFile_IsLink(t *testing.T) {
	res := IsLink("../testdata1")
	assert.Equal(t, false, res)
	res = IsLink("../testdata")
	assert.Equal(t, false, res)
}

func TestFile_TryOpen(t *testing.T) {

	type param struct {
		Path string
		Flag int
	}
	tests := []struct {
		name       string
		args       param
		want       string
		wantErr    bool
		needDelete bool
	}{
		{name: "success1",
			args: param{
				Path: "../testdata/data.txt",
				Flag: 664,
			},
			want:       "1000",
			wantErr:    false,
			needDelete: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := TryOpen(tt.args.Path, tt.args.Flag)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s is error:%s", tt.name, err.Error())
				return
			}
			if tt.needDelete {
				deletFile(tt.args.Path)
			}
		})
	}

}

func TestFile_WriteToFile(t *testing.T) {
	type param struct {
		Path string
		Text []byte
	}
	tests := []struct {
		name       string
		args       param
		want       string
		wantErr    bool
		needDelete bool
	}{
		{name: "success1",
			args: param{
				Path: "../testdata/t/data2.txt",
				Text: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			},
			want:       "1000",
			wantErr:    false,
			needDelete: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteToFile(tt.args.Path, tt.args.Text)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s is error:%s", tt.name, err.Error())
				return
			}
			if tt.needDelete {
				deletFile(tt.args.Path)
			}
		})
	}

}

func TestFile_CopyFile(t *testing.T) {
	type param struct {
		Path string
		dst  string
	}
	tests := []struct {
		name       string
		args       param
		want       string
		wantErr    bool
		needDelete bool
	}{
		{name: "error_1",
			args: param{
				Path: "../testdata/data1.txt",
				dst:  "../testdata/data.txt",
			},
			want:       "",
			wantErr:    true,
			needDelete: false,
		},
		{name: "error_2",
			args: param{
				Path: "../testdata/data.txt",
				dst:  "../testdata/data.txt",
			},
			want:       "",
			wantErr:    true,
			needDelete: false,
		},
		{name: "success",
			args: param{
				Path: "../testdata/data.txt",
				dst:  "../testdata/data1.txt",
			},
			want:       "",
			wantErr:    false,
			needDelete: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CopyFile(tt.args.Path, tt.args.dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s is error:%s", tt.name, err.Error())
				return
			}
			if tt.needDelete {
				deletFile(tt.args.dst)
			}
		})
	}

}

/*
	func TestFile_CopyDir(t *testing.T) {
		oFile := "/data.txt"
		expDir := "../testdata/t1"
		expFile := expDir + oFile
		_, err := TryOpen(expFile, 664)
		if err != nil {
			t.Errorf("CopyDir expDir error:%s", err)
			return
		}
		defer func() {
			deletFile(expFile)
		}()
		type param struct {
			Path string
			dst  string
		}
		tests := []struct {
			name       string
			args       param
			want       string
			wantErr    bool
			needDelete bool
		}{
			{name: "Err src not exist",
				args: param{
					Path: expDir + "ERROR",
					dst:  "../testdata/t2",
				},
				want:       "",
				wantErr:    true,
				needDelete: true,
			},
			{name: "Err src not dir",
				args: param{
					Path: expFile,
					dst:  "../testdata/t2",
				},
				want:       "",
				wantErr:    true,
				needDelete: true,
			},
			{name: "Err des exist",
				args: param{
					Path: expDir,
					dst:  "../testdata/t3",
				},
				want:       "",
				wantErr:    true,
				needDelete: false,
			},
			{name: "success",
				args: param{
					Path: expDir,
					dst:  "../testdata/t2",
				},
				want:       "",
				wantErr:    false,
				needDelete: true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := CopyDir(tt.args.Path, tt.args.dst)
				if (err != nil) != tt.wantErr {
					t.Errorf("%s is error:%s", tt.name, err.Error())
					return
				}
				if tt.needDelete {
					deletFile(tt.args.dst + oFile)
				}
			})
		}

}
*/
func deletFile(path string) {
	fabs, _ := filepath.Abs(path)
	os.Remove(fabs)
	os.Remove(filepath.Dir(fabs))
}
