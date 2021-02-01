package timeutil

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDate(t *testing.T) {
	date := Date()
	tm, err := ParseDate(date)

	if err != nil || FormatDate(tm) != date {
		t.Fatalf("parse date error")
	}
}

func TestParseDateTime(t *testing.T) {
	date := DateTime()
	tm, err := ParseDateTime(date)

	if err != nil || FormatDateTime(tm) != date {
		t.Fatalf("parse datetime error")
	}
}

func TestFormatDate(t *testing.T) {
	testDate := "2012-10-24"
	tm, err := ParseDate(testDate)

	if err != nil || FormatDate(tm) != testDate {
		t.Fatalf("format date error")
	}
}

func TestFormatDateTime(t *testing.T) {
	testDate := "2012-10-24 07:21:15"
	tm, err := ParseDateTime(testDate)

	if err != nil || FormatDateTime(tm) != testDate {
		t.Fatalf("format date error")
	}
}

func TestStartTime(t *testing.T) {
	testDate, _ := ParseDateTime("2012-10-24 07:21:15")
	tm := StartTime(testDate)
	startTime, _ := ParseDateTime("2012-10-24 00:00:00")

	if tm != startTime {
		t.Fatalf("start time parse error")
	}
}

func TestEndTime(t *testing.T) {
	testDate, _ := ParseDateTime("2012-10-24 07:21:15")
	tm := EndTime(testDate)
	endTime, _ := ParseDateTime("2012-10-24 23:59:59")

	if tm != endTime {
		t.Fatalf("start time parse error")
	}
}

func TestToUnix(t *testing.T) {
	testDate := "2012-10-24 07:21:15"

	if ToUnix(testDate) != 1351034475 {
		t.Fatalf("to unix error")
	}
}
func TestSplitTime(t *testing.T) {
	startTime, err := ParseDateTime("2018-08-01 00:00:00")
	if err != nil {
		t.Fatalf("parse datetime error")
	}
	endTime, err := ParseDateTime("2018-08-01 23:59:59")
	if err != nil {
		t.Fatalf("parse datetime error")
	}

	res := SplitTime(startTime, endTime, 2*time.Hour)

	// for _,v := range res{
	//	fmt.Println(v.StartTime,v.EndTime)
	// }

	l := len(res)
	if l != 12 {
		t.Errorf("time split error,want splited length 12,got %d \n", l)
	}

	if res[0].EndTime.Format("2006-01-02 15:04:05") != "2018-08-01 01:59:59" {
		t.Error("time split item error")
	}
}

func TestFormatUnix(t *testing.T) {
	testDate := "2018-11-02 13:48:00"
	ti := ToUnix(testDate)
	f := FormatUnix(ti, "2006-01-02 15:04:05")
	if FormatUnix(ti, "2006-01-02 15:04:05") != "2018-11-02 13:48:00" {
		t.Errorf("unix format date error must get %s but get %s", testDate, f)
	}
}

func TestIsZeroTime(t *testing.T) {

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{t: time.Time{}}, want: true},
		{name: "test2", args: args{t: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)}, want: true},
		{name: "test3", args: args{t: time.Date(1, 0, 0, 0, 0, 0, 0, time.UTC)}, want: true},
		{name: "test4", args: args{t: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)}, want: true},
		{name: "test5", args: args{t: time.Date(1000, 0, 0, 0, 0, 0, 0, time.UTC)}, want: true},
		{name: "test6", args: args{t: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)}, want: true},
		{name: "test7", args: args{t: time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC)}, want: false},
		{name: "test8", args: args{t: time.Now()}, want: false},
		{name: "test9", args: args{t: Zero()}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZeroTime(tt.args.t); got != tt.want {
				t.Errorf("IsZeroTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

type tt struct {
	CreatedAt Time `json:"created_at"`
}

func TestTime_UnmarshalJSON(t *testing.T) {
	var tt = &tt{}
	ett, _ := time.ParseInLocation(timeFormat, "2012-10-24 07:09:00", loc)
	err := json.Unmarshal([]byte(`{"created_at":"2012-10-24 07:09:00"}`), &tt)
	assert.NoError(t, err)
	assert.Equal(t, ett.Format(timeFormat), tt.CreatedAt.Format(timeFormat))
}

func TestTime_MarshalJSON(t *testing.T) {
	ett, _ := time.ParseInLocation(timeFormat, "2012-10-24 07:09:00", loc)
	var tt = &tt{
		CreatedAt: Time{ett},
	}
	tBytes, err := json.Marshal(tt)
	assert.NoError(t, err)
	assert.Equal(t, `{"created_at":"2012-10-24 07:09:00"}`, string(tBytes))
}

func TestMustParseDataTime(t *testing.T) {
	assert.NotPanics(t, func() {
		d := MustParseDateTime("2012-10-24 07:09:00")
		assert.Equal(t, time.Date(2012, 10, 24, 07, 9, 0, 0, loc), d)
	})

	assert.Panics(t, func() {
		MustParseDateTime("2012-10-2407:09:00")
	})
}

func TestInterval(t *testing.T) {
	type args struct {
		now   time.Time
		start string
		end   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "in",
			args: args{MustParseDateTime("2012-08-02 00:00:00"), "2012-08-01 00:00:00", "2012-09-01 00:00:00"},
			want: true,
		},
		{
			name: "out",
			args: args{MustParseDateTime("2020-08-02 00:00:00"), "2012-08-01 00:00:00", "2012-09-01 00:00:00"},
			want: false,
		},
		{
			name: "in left contains",
			args: args{MustParseDateTime("2012-08-01 00:00:00"), "2012-08-01 00:00:00", "2012-09-01 00:00:00"},
			want: true,
		},
		{
			name: "in right contains",
			args: args{MustParseDateTime("2012-09-01 00:00:00"), "2012-08-01 00:00:00", "2012-09-01 00:00:00"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Interval(tt.args.now, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("Interval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustParseDate(t *testing.T) {
	assert.NotPanics(t, func() {
		d := MustParseDate("2012-10-24")
		assert.Equal(t, time.Date(2012, 10, 24, 0, 0, 0, 0, loc), d)
	})

	assert.Panics(t, func() {
		MustParseDate("2012-10-66")
	})
}
