package hashing

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"sort"
)

func encode(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := k + "="
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(prefix)
			buf.WriteString(v)
		}
	}
	return buf.String()
}

func Sign(values url.Values, app_secret string) (string) {
	for key, _ := range values {
		if values.Get(key) == "" || key == "sign" {
			values.Del(key)
		}
	}
	return Md5(encode(values) + app_secret)
}

func Md5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}