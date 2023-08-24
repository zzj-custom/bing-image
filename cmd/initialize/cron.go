package initialize

import (
	"github.com/robfig/cron/v3"
	"image/internal/global"
	"image/internal/task"
)

func InitCron() {
	if global.GVA_CONFIG.Cron.Start {
		// 添加定时任务
		var option []cron.Option
		if global.GVA_CONFIG.Cron.WithSeconds {
			option = append(option, cron.WithSeconds())
		}

		// 定时拉取必应每日图片
		bingImagesTask := task.NewCrawlBingImage()
		_, _ = global.GVA_Timer.AddTaskByJob(bingImagesTask.Name(), bingImagesTask.Spec(), bingImagesTask, option...)
	}
}
