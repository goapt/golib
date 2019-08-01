package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Request struct {
	Msgtype string                 `json:"msgtype"`
	Text    map[string]interface{} `json:"text"`
}

type Response struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type MarkdownRequest struct {
	Msgtype  string                 `json:"msgtype"`
	Markdown map[string]interface{} `json:"markdown"`
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
	we := &MarkdownRequest{
		Msgtype: "markdown",
		Markdown: map[string]interface{}{
			"content": md,
		},
	}

	if len(at) > 0 {
		we.Markdown["mentioned_mobile_list"] = at
	}

	buf, err := json.Marshal(we)
	if err != nil {
		return err
	}

	return r.call(buf)
}

func (r *Robot) call(buf []byte) error {
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + r.Token
	resp, err := r.postJson(url, buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if data, err := ioutil.ReadAll(resp.Body); err == nil {
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
	we := &Request{
		Msgtype: "text",
		Text: map[string]interface{}{
			"content": content,
		},
	}

	if len(at) > 0 {
		we.Text["mentioned_mobile_list"] = at
	}

	buf, err := json.Marshal(we)
	if err != nil {
		return err
	}

	return r.call(buf)
}

func (r *Robot) CardMessage(title, text string, btns []map[string]string) error {
	return nil
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
