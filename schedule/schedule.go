package schedule

import (
	"time"
)

func ScheduleHourly(task func()) {
	for {
		task() // 执行指定的函数
		time.Sleep(time.Minute * 15)
	}
}

func ScheduleHourly10(task func()) {
	for {
		task() // 执行指定的函数
		time.Sleep(time.Minute * 10)
	}
}
