package model

import (
	"testing"
)

type t_users struct {
	UserId       int    `json:"user_id" xorm:"pk"`
	Type         int    `json:"type" `
	AuthMode     int    `json:"auth_mode" `
	MerchantName string `json:"merchant_name" `
	UserName     string `json:"user_name" `
	Password     string `json:"password" `
	Status       int    `json:"status" `
	ParentUserId int    `json:"parent_user_id" `
	ApiKey       string `json:"api_key" `
	ApiSecret    string `json:"api_secret" `
	NotifySecret string `json:"notify_secret" `
	Category     int    `json:"category" `
	Timezone     string `json:"timezone" `
	Currency     string `json:"currency" `
	Source       int    `json:"source" `
	SourceId     int    `json:"source_id" `
	Lang         string `json:"lang" `
	CreateTime   Time   `xorm:"created" json:"create_time" `
	UpdateTime   Time   `xorm:"updated" json:"update_time" `
}

var m = map[string]string{
	"user_id":        "4",
	"type":           "3",
	"auth_mode":      "1",
	"merchant_name":  "test",
	"user_name":      "test",
	"password":       "xxxxxxx",
	"status":         "1",
	"parent_user_id": "3",
	"api_key":        "xxxx",
	"api_secret":     "xxxx",
	"notify_secret":  "xxx",
	"category":       "0",
	"timezone":       "Asia/Shanghai",
	"currency":       "CNY",
	"source":         "1",
	"source_id":      "0",
	"lang":           "zh-CN",
	"create_time":    "2015-03-18 18:20:28",
	"update_time":    "2017-09-20 10:29:59",
}

func BenchmarkScanStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := &t_users{}
		ScanStruct(m, user)
	}
}

//func BenchmarkScanStruct2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		user := &t_users{}
//		ScanStruct2(m, user)
//	}
//}

func TestScanStruct(t *testing.T) {
	user := &t_users{}
	err := ScanStruct(m, user)

	if err != nil {
		t.Error(err)
	}

	if user.UserId != 4 {
		t.Error("Parse user_id error:", user.UserId)
	}

	if user.CreateTime.String() != "2015-03-18 18:20:28" {
		t.Error("Parse user_id error:", user.CreateTime)
	}

}