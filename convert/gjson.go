package convert

import "github.com/tidwall/gjson"

type JsonTo string

func (j JsonTo) MapString() map[string]string {
	r := gjson.Parse(string(j))
	return toSSMap(r)
}

type JsonBytesTo []byte

func (j JsonBytesTo) MapString() map[string]string {
	r := gjson.ParseBytes([]byte(j))
	return toSSMap(r)
}

func toSSMap(r gjson.Result) map[string]string {
	rs := make(map[string]string)
	if !r.Exists() {
		return rs
	}
	for k, v := range r.Map() {
		rs[k] = v.String()
	}
	return rs
}
