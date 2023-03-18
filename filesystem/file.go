package filesystem

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// IsWritable path is writable
func IsWritable(path string) bool {
	err := syscall.Access(path, syscall.O_RDWR)
	return err == nil
}

// IsReadable path is readable
func IsReadable(path string) bool {
	err := syscall.Access(path, syscall.O_RDONLY)
	return err == nil
}

// IsExecutable path is executable
func IsExecutable(file string) bool {
	info, err := os.Stat(file)
	return err == nil && info.Mode().IsRegular() && (info.Mode()&0111) != 0
}

// IsExists path is exists
func IsExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsDir determines whether the specified path is a directory.
func IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if nil != err {
		return false
	}

	return fio.IsDir()
}

// IsLink path is symlink
func IsLink(path string) bool {
	f, err := os.Lstat(path)
	if err != nil {
		return false
	}

	return f.Mode()&os.ModeSymlink == os.ModeSymlink
}

// ReadLine read file line output channel
func ReadLine(filePth string) (chan []byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}

	lc := make(chan []byte)

	go func() {
		defer f.Close()

		br := bufio.NewReader(f)
		for {
			// 此处不能使用br.ReadLine() 这个方法存在问题，对并发读取的时候会重复
			a, c := br.ReadString('\n')
			if c == io.EOF {
				close(lc)
				break
			}

			str := strings.TrimRight(a, "\r\n")

			lc <- []byte(str)
		}
	}()

	return lc, nil
}

// WriteToFile write byte to file
func WriteToFile(filename string, text []byte) error {
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.Write(text); err != nil {
		return err
	}
	return nil
}

// CopyFile copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file. The file mode will be copied from the source and
// the copied data is synced/flushed to stable storage.
func CopyFile(src, dst string) (err error) {
	srcAbs, _ := filepath.Abs(src)
	dstAbs, _ := filepath.Abs(dst)
	if srcAbs == dstAbs {
		err = errors.New("Cannot use the same file")
		return
	}
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}

// CopyDir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
// Symlinks are ignored and skipped.
func CopyDir(src string, dst string) error {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err == nil {
		return fmt.Errorf("destination already exists")
	}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			info, err := entry.Info()
			if err != nil {
				return err
			}
			// Skip symlinks.
			if info.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// TryOpen 打开一个文件，如果文件不存在，则尝试创建
func TryOpen(path string, flag int) (*os.File, error) {
	fabs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(fabs, flag, os.ModePerm)
	if os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(fabs), os.ModePerm)
		if err != nil {
			return nil, err
		}
		return os.OpenFile(fabs, flag, os.ModePerm)
	}
	return f, err
}
