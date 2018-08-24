package timeutil

import "time"

type TimeRange struct {
	StartTime string
	EndTime   string
}

func Date() string {
	return time.Now().Format("2006-01-02")
}

func DateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func ParseDate(tstr string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation("2006-01-02", tstr, loc)
}

func ParseDateTime(tstr string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation("2006-01-02 15:04:05", tstr, loc)
}

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func ToUnix(tstr string) int64 {
	t, err := ParseDateTime(tstr)

	if err != nil {
		return 0
	}
	return t.Unix()
}

func SplitTime(startTime, endTime string, duration time.Duration) ([]*TimeRange, error) {
	var (
		err error
		s   time.Time
		e   time.Time
		t   time.Time
		r   = make([]*TimeRange, 0)
	)
	if s, err = ParseDateTime(startTime); err != nil {
		return nil, err
	}
	if e, err = ParseDateTime(endTime); err != nil {
		return nil, err
	}

	for s.Before(e) {
		tr := &TimeRange{}
		tr.StartTime = FormatDateTime(s)

		t = s.Add(duration - time.Second)
		if t.After(e) {
			t = e
		}
		tr.EndTime = FormatDateTime(t)
		s = s.Add(duration)
		r = append(r, tr)
	}

	return r, nil
}
