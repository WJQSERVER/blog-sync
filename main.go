package main

import (
	"fmt"
	"log"
	"time"

	"blog-sync/build"
	"blog-sync/config"
	"blog-sync/download"
	"blog-sync/logger"
	"blog-sync/schedule"
)

var cfg *config.Config
var logw = logger.Logw

func loadConfig() {
	var err error
	// 初始化配置
	cfg, err = config.LoadConfig("/root/data/blog-sync/config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Printf("Loaded config: %v\n", cfg)
}

func setupLogger() {
	// 初始化日志模块
	var err error
	err = logger.Init(cfg.LogFilePath) // 传递日志文件路径
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	logw("Logger initialized")
	logw("Init Completed")
}

func setupSchedule() {
	schedule.Init(cfg.CycleInterval)
}

func init() {
	loadConfig()
	setupLogger()
	setupSchedule()
}

func main() {
	defer logger.Close() // 确保在退出时关闭日志文件
	go schedule.Schedule(func() {
		err := download.DownloadFile(cfg.DownloadUrl, cfg.Username, cfg.Password, cfg.SavePath, cfg.Unzipdir)
		if err != nil {
			logw("下载文件时出错: %v", err) // 处理错误
		}
	})

	go schedule.Schedule(func() {
		var sleep int
		sleep = 5
		logw("开始执行hugo构建任务，等待%d分钟", sleep)
		time.Sleep(time.Duration(sleep) * time.Minute)
		err := build.Build(cfg.Unzipdir, cfg.BaseURL)
		if err != nil {
			logw("Hugo执行构建任务时出错: %v", err) // 处理错误
		}
	})
	select {}
}
