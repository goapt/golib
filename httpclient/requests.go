package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/verystar/golib/logger"
	"github.com/verystar/golib/monitor"
)

const (
	FORM_URLENCODED = "application/x-www-form-urlencoded"
	JSON            = "application/json"
	XML             = "application/xml"
)

type HttpClient struct {
	*http.Client
}

func NewClient(options ...func(*HttpClient)) *HttpClient {
	httpclient := &HttpClient{
		&http.Client{},
	}
	httpclient.Client.Timeout = 5 * time.Second
	// Apply options in the parameters to request.
	for _, option := range options {
		option(httpclient)
	}

	return httpclient
}

func (h *HttpClient) Post(url string, contentType string, body io.Reader, options ... func(r *http.Request)) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)

	for _, option := range options {
		option(req)
	}
	return h.Do(req)
}

func (h *HttpClient) Get(url string, options ... func(r *http.Request)) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for _, option := range options {
		option(req)
	}
	return h.Do(req)
}

func (h *HttpClient) Do(req *http.Request) (resp *http.Response, err error) {
	start := time.Now()
	rep, err := h.Client.Do(req)

	monitorPush(start, req.URL.String(), rep, err)
	return rep, err
}

func parseErr(str string) string {
	errs := make(map[string]string)
	errs["Client.Timeout"] = "HTTP请求超时(28)"
	errs["no such host"] = "DNS解析失败(6)"
	errs["unsupported protocol scheme"] = "网址格式不正确(3)"

	for k, v := range errs {
		if strings.Index(str, k) != -1 {
			return v
		}
	}

	return "错误"
}

func monitorPush(start time.Time, urlStr string, req *http.Response, err error) {
	end := time.Now()
	urls, _ := url.Parse(urlStr)
	stat_url := urls.Host + urls.Path
	diff_time_str := fmt.Sprintf("%.6f", end.Sub(start).Seconds())
	diff_time, _ := strconv.ParseFloat(diff_time_str, 64)

	monitor.Stat(1, "CURL接口效率", monitor.FormatTime(diff_time), stat_url)

	if err == nil {
		if req.StatusCode == 200 {
			monitor.Stat(1, "CURL请求", "成功", urls.Host)
		}

		monitor.Stat(1, "CURL状态", strconv.Itoa(req.StatusCode), urls.Host)
	} else {
		monitor.Stat(1, "CURL请求", parseErr(err.Error()), urls.Host)

		logger.Info("[CURL ERROR]%s%+v", err.Error(), map[string]string{
			"Url":         urlStr,
			"RequestTime": diff_time_str,
			"ErrorInfo":   err.Error(),
		})

		if req != nil {
			monitor.Stat(1, "CURL状态", strconv.Itoa(req.StatusCode), urls.Host)
		}
	}
}