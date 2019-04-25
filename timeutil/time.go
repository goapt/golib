package timeutil

import (
	"time"
)

type TimeRange struct {
	StartTime time.Time
	EndTime   time.Time
}

var loc, _ = time.LoadLocation("Asia/Shanghai")

func Date() string {
	return time.Now().In(loc).Format("2006-01-02")
}

func DateTime() string {
	return time.Now().In(loc).Format("2006-01-02 15:04:05")
}

func ParseDate(tstr string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", tstr, loc)
}

func ParseDateTime(tstr string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	if IsRFC3339(tstr) {
		layout = time.RFC3339
	}

	return time.ParseInLocation(layout, tstr, loc)
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

func FormatDate(t time.Time) string {
	return t.In(loc).Format("2006-01-02")
}

func FormatDateTime(t time.Time) string {
	return t.In(loc).Format("2006-01-02 15:04:05")
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
