package taskscheduler

import (
	"time"

	"github.com/go-co-op/gocron"
)

func ConfigTaskScheduler() *gocron.Scheduler {
	schedule := gocron.NewScheduler(time.UTC)

	return schedule
}