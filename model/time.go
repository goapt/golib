package model

import (
	"time"
)

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	return t.marshal(true)
}

func (t Time)MarshalBinary() ([]byte, error) {
	return t.marshal(false)
}

func (t Time)marshal(is_append bool) ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	if is_append {
		b = append(b, '"')
	}
	b = time.Time(t).AppendFormat(b, timeFormart)
	if is_append {
		b = append(b, '"')
	}
	return b, nil
}
func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}

func (t Time) Format(layout string) string {
	return time.Time(t).Format(layout)
}

func (t Time) ToTime() time.Time {
	return time.Time(t)
}

func Now() Time {
	return Time(time.Now())
}
