package Controller

import (
	"time"

	"github.com/robfig/cron/v3"
)

// SchedulerController merupakan kontroler untuk menjadwalkan tugas secara berkala
type SchedulerController struct {
	Interval time.Duration
}

// NewSchedulerController membuat instance baru dari SchedulerController
func NewSchedulerController(interval time.Duration) *SchedulerController {
	return &SchedulerController{
		Interval: interval,
	}
}

func (sc *SchedulerController) ScheduleTask(task func()) {
	cronJob := cron.New()
	_, _ = cronJob.AddFunc("0 0 * * *", task)
	cronJob.Start()
}
