package utils

import "time"

func GetTimeout(timeout ...time.Duration) time.Duration {
	var t time.Duration
	if len(timeout) > 0 {
		t = timeout[0]
	} else {
		t = 60 * time.Second // 默认超时时间为 60 秒
	}
	return t
}
