package schedule

import (
	"blog-sync/logger"
	"time"
)

var waittime int
var logw = logger.Logw

func Init(roundtime int) {
	waittime = roundtime
}

func Schedule(task func()) {
	for {
		task() // 执行指定的函数
		logw("schedule task done, wait for %d minutes", waittime)
		time.Sleep(time.Minute * time.Duration(waittime))
	}
}
