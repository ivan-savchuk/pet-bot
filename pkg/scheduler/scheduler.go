package scheduler

import (
	"log"
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
	_, err := s.scheduler.Cron(crontab).Do(Job)
	if err != nil {
		log.Panicf("\nCould not set new job for provide crontab '%v'", crontab)
	}
}

func (s *Scheduler) Start() {
	s.scheduler.StartAsync()
}
