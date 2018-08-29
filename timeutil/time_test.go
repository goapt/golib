package timeutil

import (
	"testing"
	"time"
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

	//for _,v := range res{
	//	fmt.Println(v.StartTime,v.EndTime)
	//}

	l := len(res)
	if l != 12 {
		t.Errorf("time split error,want splited length 12,got %d \n", l)
	}

	if res[0].EndTime.Format("2006-01-02 15:04:05") != "2018-08-01 01:59:59" {
		t.Error("time split item error")
	}
}
