package debug

import (
	"time"
	"encoding/json"
	"bytes"
	"fmt"
	"github.com/verystar/golib/color"
	"strings"
	"strconv"
	"path/filepath"
	"os"
)

type DebugTagData struct {
	Key     string
	Data    interface{}
	Stack   callStack
	Current string
}

type DebugTag struct {
	t    time.Time
	data []DebugTagData
}

func NewDebugTag(options ...func(*DebugTag)) *DebugTag {
	debug := &DebugTag{}

	for _, option := range options {
		option(debug)
	}

	debug.Start()
	return debug
}

func (this *DebugTag) Start() {
	if debugFlag == "off" {
		return
	}
	this.t = time.Now()
}

func (this *DebugTag) Tag(key string, data ...interface{}) {
	if debugFlag == "off" {
		return
	}

	st := Callstack(2)
	t := time.Now().Sub(this.t).String()

	if printTag == "" || strings.Contains(key, printTag) {
		fmt.Println(color.Blue("[Debug Tag]("+t+")") + " -------------------------> " + key + " <-------------------------")
		fmt.Println(color.Green("File:" + st.File + ", Func:" + st.Func + ", Line:" + strconv.Itoa(st.LineNo)))
		if len(data) > 0 {
			format := strings.Repeat("===> %v\n", len(data))
			fmt.Println(color.Yellow(format, data...))
		}
	}

	this.data = append(this.data, DebugTagData{
		Key:     key,
		Data:    data,
		Stack:   st,
		Current: t,
	})
}

func (this *DebugTag) GetTagData() []DebugTagData {
	return this.data
}

func (this *DebugTag) Save(dir string, format string, prefix ...string) error {
	pre := ""
	if len(prefix) > 0 {
		pre = prefix[0] + "_"
	}
	if debugFlag == "off" {
		return nil
	}

	now := time.Now()
	s := now.Format(format)
	filename := strings.TrimRight(savePath, "/") + "/" + dir + "/" + pre + s + ".log"
	//buf , err := json.Marshal(this.data)
	buf, err := json.MarshalIndent(this.data, "", "    ")
	if err != nil {
		return err
	}
	buffer := bytes.NewBufferString(fmt.Sprintf("\n[%v]\n", now.String()))
	buffer.Write(buf)
	buffer.WriteString("\n\n")
	return writeToFile(filename, buffer.Bytes())
}

func writeToFile(filename string, text []byte) error {
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.Write(text); err != nil {
		return err
	}
	return nil
}

func (this *DebugTag) SaveToSecond(dir string, prefix ...string) error {
	return this.Save(dir, "2006-01-02-15-04-05", prefix...)
}

func (this *DebugTag) SaveToMinute(dir string, prefix ...string) error {
	return this.Save(dir, "2006-01-02-15-04", prefix...)
}

func (this *DebugTag) SaveToHour(dir string, prefix ...string) error {
	return this.Save(dir, "2006-01-02-15", prefix...)
}

func (this *DebugTag) SaveToDay(dir string, prefix ...string) error {
	return this.Save(dir, "2006-01-02", prefix...)
}
