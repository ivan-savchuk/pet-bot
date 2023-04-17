package scheduler

import (
	"time"

	gocron "github.com/go-co-op/gocron"
)

type Scheduler struct {
	scheduler *gocron.Scheduler
}

func (s *Scheduler) SetNewScheduler(timezone *time.Location) *gocron.Scheduler {
	if s.scheduler == nil {
		s.scheduler = gocron.NewScheduler(timezone)
	}
	return s.scheduler
}

func (s *Scheduler) AddNewJob(crontab string, Job func()) {
	s.scheduler.Cron(crontab).Do(Job)
}

func (s *Scheduler) Start() {
	s.scheduler.StartAsync()
}
