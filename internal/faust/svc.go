package faust

import (
	"errors"
	"faust/internal/pkg/config"
	"faust/internal/pkg/notify"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	svc *Server
)

type Server struct {
	bot     *tgbotapi.BotAPI
	cfg     *config.Config
	updater tgbotapi.UpdatesChannel
}

func NewFaustSvc() (*Server, error) {
	cfg := config.GetConfig()
	if "" == cfg.Token {
		return nil, errors.New("token not configured")
	}

	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if nil != err {
		return nil, err
	}

	bot.Debug = cfg.DebugMode

	svc := Server{
		bot: bot,
		cfg: cfg,
	}

	svc.RegisterSignal()
	return &svc, nil
}

func (s Server) Serve() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = s.cfg.Timeout

	updater, err := s.bot.GetUpdatesChan(u)
	if nil != err {
		return err
	}
	s.updater = updater

	for update := range s.updater {
		if nil == update.Message {
			continue
		}

		c, err := NewContext(s.bot, &update)
		if nil != err {
			return err
		}

		Dispatch(c)
	}

	return nil
}

func (s Server) RegisterSignal() {
	notify.SigBeep.Connect(HandleBeepNotify)
}

func GetSvc() *Server {
	if nil != svc {
		return svc
	}

	svc, err := NewFaustSvc()
	if nil != err {
		panic(err)
	}

	return svc
}
