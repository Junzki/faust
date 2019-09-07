package faust

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"sync"
)


type SubscriberMap map[int64] *tgbotapi.Chat

type SubscriberContainer struct {
	m     SubscriberMap
	mutex sync.RWMutex
}

func (s *SubscriberContainer) Subscribe(c *tgbotapi.Chat) error {
	if nil == c {
		return errors.New("chat is nil")
	}

	s.mutex.Lock()

	id := c.ID
	_, existed := s.m[id]
	if existed {
		s.mutex.Unlock()
		return errors.New("already subscribed")
	}

	s.m[id] = c
	s.mutex.Unlock()
	return nil
}

func (s *SubscriberContainer) Unsubscribe(id int64) error {
	s.mutex.Lock()

	_, existed := s.m[id]
	if ! existed {
		s.mutex.Unlock()
		return errors.New("not subscribed")
	}

	delete(s.m, id)

	s.mutex.Unlock()
	return nil
}

func (s *SubscriberContainer) Broadcast(bot *tgbotapi.BotAPI, text string) error {
	s.mutex.RLock()

	for _, c := range s.m {
		rep := tgbotapi.NewMessage(c.ID, text)
		_, err := bot.Send(rep)
		if nil != err {
			log.Warn(err)
		}
	}

	return nil
}


var (
	BeepSubscriber = SubscriberContainer{
		m: make(SubscriberMap),
	})
