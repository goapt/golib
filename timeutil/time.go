package timeutil

import (
	"fmt"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	dateFormat = "2006-01-02"
	timeZone   = "Asia/Shanghai"
)

type TimeRange struct {
	StartTime time.Time
	EndTime   time.Time
}

var loc, _ = time.LoadLocation(timeZone)

func Date() string {
	return time.Now().In(loc).Format(dateFormat)
}

func DateTime() string {
	return time.Now().In(loc).Format(timeFormat)
}

func ParseDate(tstr string) (time.Time, error) {
	return time.ParseInLocation(dateFormat, tstr, loc)
}

func ParseDateTime(tstr string) (time.Time, error) {
	layout := timeFormat
	if IsRFC3339(tstr) {
		layout = time.RFC3339
	}

	return time.ParseInLocation(layout, tstr, loc)
}

func MustParseDate(t string) time.Time {
	v, err := ParseDate(t)
	if err != nil {
		panic(err)
	}
	return v
}

func MustParseDateTime(t string) time.Time {
	v, err := ParseDateTime(t)
	if err != nil {
		panic(err)
	}
	return v
}

// 自定义格式解析时间
func ParseDateTimeInLayout(tstr, layout string) (time.Time, error) {
	return time.ParseInLocation(layout, tstr, loc)
}

//
func MustParseDateTimeInLayout(tstr, layout string) time.Time {
	v, _ := ParseDateTimeInLayout(tstr, layout)
	return v
}

// IsRFC3339 check if string is valid timestamp value according to RFC3339
func IsRFC3339(str string) bool {
	return IsTime(str, time.RFC3339)
}

// IsTime check if string is valid according to given format
func IsTime(str string, format string) bool {
	_, err := time.Parse(format, str)
	return err == nil
}

func StartTime(t time.Time) time.Time {
	t, _ = ParseDateTime(FormatDate(t) + " 00:00:00")
	return t
}

func EndTime(t time.Time) time.Time {
	t, _ = ParseDateTime(FormatDate(t) + " 23:59:59")
	return t
}

func FormatDate(t time.Time) string {
	return t.In(loc).Format(dateFormat)
}

func FormatDateTime(t time.Time) string {
	return t.In(loc).Format(timeFormat)
}

func FormatUnix(t int64, layout string) string {
	return time.Unix(t, 0).In(loc).Format(layout)
}

func ToUnix(tstr string) int64 {
	t, err := ParseDateTime(tstr)

	if err != nil {
		return 0
	}
	return t.Unix()
}

// Interval 判断时间是否在start,end区间，包含start,end
func Interval(now time.Time, start, end string) bool {
	st := MustParseDateTime(start)
	et := MustParseDateTime(end)
	return now.Unix() >= st.Unix() && now.Unix() <= et.Unix()
}

func Zero() time.Time {
	t, _ := time.Parse(timeFormat, "0000-00-00 00:00:00")
	return t
}

// IsZeroTime 判断时间是否为零值
// 比较常见的零值时间为：
// 0000-00-00 00:00:00
// 0000-00-01 00:00:00
// 0001-01-01 00:00:00
// 1000-01-01 00:00:00
// 1970-01-01 00:00:00
func IsZeroTime(t time.Time) bool {
	// 比Unix纪元时间小的，都是零值
	return time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC).Sub(t) >= 0
}

func SplitTime(startTime, endTime time.Time, duration time.Duration) []*TimeRange {
	var (
		t time.Time
		r = make([]*TimeRange, 0)
	)

	for startTime.Before(endTime) {
		tr := &TimeRange{}
		tr.StartTime = startTime

		t = startTime.Add(duration - time.Second)
		if t.After(endTime) {
			t = endTime
		}
		tr.EndTime = t
		startTime = startTime.Add(duration)
		r = append(r, tr)
	}

	return r
}

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Format(timeFormat))), nil
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	t.Time, err = time.ParseInLocation(fmt.Sprintf(`"%s"`, timeFormat), string(data), loc)
	return
}
