package ding

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var DING_TALK_TOKEN string

type dingTalkRequest struct {
	Msgtype string                 `json:"msgtype"`
	Text    map[string]string      `json:"text"`
	At      map[string]interface{} `json:"at,omitempty"`
}

type dingTalkResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type dingTalkMarkdownRequest struct {
	Msgtype  string                 `json:"msgtype"`
	Markdown map[string]string      `json:"markdown"`
	At       map[string]interface{} `json:"at,omitempty"`
}

func AlarmMarkdown(md string, at ...string) error {
	dingtalk := &dingTalkMarkdownRequest{
		Msgtype: "markdown",
		Markdown: map[string]string{
			"title": md[0:20],
			"text":  md,
		},
	}

	if len(at) > 0 {
		dingtalk.At = map[string]interface{}{
			"atMobiles": at,
		}
	}

	buf, err := json.Marshal(dingtalk)
	if err != nil {
		return err
	}

	fmt.Println(string(buf))

	return call(buf)
}

func call(buf []byte) error {
	url := "https://oapi.dingtalk.com/robot/send?access_token=" + DING_TALK_TOKEN
	resp, err := postJson(url, buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if data, err := ioutil.ReadAll(resp.Body); err == nil {
		ret := &dingTalkResponse{}
		err := json.Unmarshal(data, ret)
		if err != nil {
			return err
		}

		if ret.Errcode != 0 {
			return errors.New("ding response error:" + ret.Errmsg + "[" + strconv.Itoa(ret.Errcode) + "]")
		}
	}

	if err != nil {
		return err
	}
	return nil
}

func Alarm(content string, at ...string) error {
	dingtalk := &dingTalkRequest{
		Msgtype: "text",
		Text: map[string]string{
			"content": content,
		},
	}

	if len(at) > 0 {
		dingtalk.At = map[string]interface{}{
			"atMobiles": at,
		}
	}

	buf, err := json.Marshal(dingtalk)
	if err != nil {
		return err
	}

	return call(buf)
}

func postJson(url string, data []byte) (*http.Response, error) {
	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	client := &http.Client{}
	client.Timeout = 5 * time.Second

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
