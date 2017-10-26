package model

import (
	"testing"
	"time"
	"encoding/json"
)

type t_time_json struct {
	UserName   string `json:"user_name"`
	CreateTime Time   `json:"create_time"`
}

func TestTime_MarshalJSON(t *testing.T) {

	tj := &t_time_json{
		UserName:   "test",
		CreateTime: Time(time.Now()),
	}

	j, err := json.Marshal(tj)

	if err != nil {
		t.Error(err)
	}

	tj2 := &t_time_json{}
	err = json.Unmarshal(j, tj2)

	if err != nil {
		t.Error(err)
	}

	if tj.CreateTime.String() != tj2.CreateTime.String() {
		t.Error("Time cannot be reduced:", tj.CreateTime, tj2.CreateTime)
	}
}
