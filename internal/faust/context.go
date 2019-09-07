package faust

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

func NewContext(b *tgbotapi.BotAPI, u *tgbotapi.Update) (*Context, error) {
	if nil == u || nil == b {
		return nil, errors.New("required arg is nil")
	}

	c := Context{
		Update: u,
		bot:    b,
	}

	return &c, nil
}

type Context struct {
	Update *tgbotapi.Update
	bot    *tgbotapi.BotAPI
}

// GetBot returns referenced BotAPI instance.
func (c *Context) GetBot() *tgbotapi.BotAPI {
	if nil == c.bot {
		panic("bad bot api reference")
	}

	return c.bot
}

/************************************/
/***** GOLANG.ORG/X/NET/CONTEXT *****/
/************************************/

func (c *Context) Err() error {
	return nil
}

func (c *Context) Done() <-chan struct{} {
	return nil
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c *Context) Value(key interface{}) interface{} {
	return nil
}
