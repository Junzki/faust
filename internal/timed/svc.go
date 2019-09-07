package timed

import (
	"faust/internal/pkg/notify"
	"time"
)

const (
	DefaultTimeInterval = time.Duration(60 * time.Second)
)


type Service struct {
	interval	time.Duration
	timer       *time.Timer
}

func NewService(i time.Duration) (*Service, error) {
	if 0 >= i {
		i = DefaultTimeInterval
	}

	svc := Service{
		interval: i,
	}

	return &svc, nil
}


func (s *Service) Serve() error {
	for {
		s.timer = time.NewTimer(s.interval)
		<-s.timer.C  // Will block current goroutine.

		notify.SigBeep.SendAsync(nil)
	}

	return nil
}


