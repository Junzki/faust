package faust

import (
	"faust/internal/pkg/notify"
	"fmt"
	"time"
)

const (
	DefaultTimeFormat = "Mon Jan 2 15:04:05 MST 2006"
	beepNotifyCap     = 100
)

var beepNotifyChan = make(chan time.Time, beepNotifyCap)

func HandleBeepNotify(_ notify.ISignal, _ ...interface{}) {
	now := time.Now()
	beepNotifyChan <- now

	text := fmt.Sprintf("Current time: %s.", now.Format(DefaultTimeFormat))

	svc := GetSvc()

	_ = BeepSubscriber.Broadcast(svc.bot, text)
}
