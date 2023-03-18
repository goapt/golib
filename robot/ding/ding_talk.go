package ding

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Request struct {
	Msgtype string                 `json:"msgtype"`
	Text    map[string]string      `json:"text"`
	At      map[string]interface{} `json:"at,omitempty"`
}

type Response struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type MarkdownRequest struct {
	Msgtype  string                 `json:"msgtype"`
	Markdown map[string]string      `json:"markdown"`
	At       map[string]interface{} `json:"at,omitempty"`
}

type CardRequest struct {
	Msgtype    string                 `json:"msgtype"`
	ActionCard map[string]interface{} `json:"actionCard"`
}

type Robot struct {
	Token string
}

func NewRobot() *Robot {
	return &Robot{}
}

func (r *Robot) SetToken(token string) {
	r.Token = token
}

func (r *Robot) MarkdownMessage(md string, at ...string) error {
	dingtalk := &MarkdownRequest{
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

	return r.call(buf)
}

func (r *Robot) call(buf []byte) error {
	url := "https://oapi.dingtalk.com/robot/send?access_token=" + r.Token
	resp, err := r.postJson(url, buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if data, err := io.ReadAll(resp.Body); err == nil {
		ret := &Response{}
		err := json.Unmarshal(data, ret)
		if err != nil {
			return err
		}

		if ret.Errcode != 0 {
			return errors.New("ding response error:" + ret.Errmsg + "[" + strconv.Itoa(ret.Errcode) + "]")
		}
	}

	return nil
}

func (r *Robot) Message(content string, at ...string) error {
	dingtalk := &Request{
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

	return r.call(buf)
}

func (r *Robot) CardMessage(title, text string, btns []map[string]string) error {
	dingtalk := &CardRequest{
		Msgtype: "actionCard",
		ActionCard: map[string]interface{}{
			"title":          title,
			"text":           text,
			"hideAvatar":     0,
			"btnOrientation": 1,
			"btns":           btns,
		},
	}

	buf, err := json.Marshal(dingtalk)
	if err != nil {
		return err
	}

	return r.call(buf)
}

func (r *Robot) postJson(url string, data []byte) (*http.Response, error) {
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
