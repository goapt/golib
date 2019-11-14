package number

import (
	"fmt"
	"time"
)

var (
	genRing *Ring
)

func init() {
	genRing = NewRing(10000, 99999)
}

// Unique 生成20位的唯一编号，每微秒可产生90W
func Unique(prefix string) string {
	now := time.Now()
	sec := now.Unix() % 86400              // 今天的秒数
	msec := (now.UnixNano() / 1e6) % 10000 // 1/100微秒
	// sn(20) = 前缀(0) + 日期(6) + 当日秒(5) + 微秒(4) + 随机循环数(5)
	return fmt.Sprintf("%s%s%05d%05d%04d", prefix, now.Format("060102"), sec, genRing.Next(), msec)
}
