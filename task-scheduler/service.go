package taskscheduler

import "log"

func task() {
	log.Println("this is simple task")
}

func NewSchedule() {
	newSchedule := ConfigTaskScheduler()

	newSchedule.Every(10).Seconds().Do(task)
	
	newSchedule.StartAsync()
}