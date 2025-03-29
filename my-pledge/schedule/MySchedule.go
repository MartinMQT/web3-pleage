package schedule

import (
	"log"
	"time"
)

func TestSchedule() {
	// 定义要执行的任务
	task := func() {
		log.Println("Running scheduled task at", time.Now())
	}
	go func() {
		for {
			// 执行任务
			task()
			// 等待10分钟
			time.Sleep(10 * time.Second)
		}
	}()
}
