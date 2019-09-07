package faust

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

func Dispatch(c *Context) {
	m := c.Update.Message


	if m.IsCommand() {
		handleCommand(c)
		return
	}

	return
}


const (
	CommandAloha		   = "aloha"
	CommandSubscribeBeep   = "set_beep"
	CommandUnsubscribeBeep = "unset_beep"
)

func handleCommand(c *Context) {
	rep := tgbotapi.NewMessage(c.Update.Message.Chat.ID, "")
	op := strings.ToLower(c.Update.Message.Command())

	switch op {
	case CommandAloha:
		rep.Text = fmt.Sprintf("Aloha %s", c.Update.Message.From.FirstName)
	case CommandSubscribeBeep, CommandUnsubscribeBeep:
		handleBeepSubscribe(c, op)


	default:
		return
	}


	_, _ = c.GetBot().Send(rep)
}


func handleBeepSubscribe(c *Context, op string) {
	var err error = nil
	if CommandSubscribeBeep == op {
		err = BeepSubscriber.Subscribe(c.Update.Message.Chat)
	} else {
		err = BeepSubscriber.Unsubscribe(c.Update.Message.Chat.ID)
	}

	rep := tgbotapi.NewMessage(c.Update.Message.Chat.ID, "")

	var text string
	if nil != err {
		text = fmt.Sprintf("%s failed, error: %s.", op, err.Error())
	} else {
		text = fmt.Sprintf("%s succeed.", op)
	}

	rep.Text = text

	_, _ = c.GetBot().Send(rep)
}


