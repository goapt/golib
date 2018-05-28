package convert

import "github.com/tidwall/gjson"

type JsonTo string

func (this JsonTo) MapString() map[string]string {
	r := gjson.Parse(string(this))
	return toSSMap(r)
}

type JsonBytesTo []byte

func (this JsonBytesTo) MapString() map[string]string {
	r := gjson.ParseBytes([]byte(this))
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
