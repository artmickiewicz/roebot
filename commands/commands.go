package commands

import (
	s "19u4n4/roebot/state"
	"strings"

	"robpike.io/filter"

	t "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Transition func(*t.Message) Handler

var DefaultTransition = func(m *t.Message) Handler {
	return Zero{Message: m}
}

type Handler interface {
	Handle() (transitTo Transition, reply string, sync bool)
}

type Zero struct{ Message *t.Message }

func (cmd Zero) Handle() (transitTo Transition, reply string, sync bool) {
	transitTo = DefaultTransition
	reply = ""
	sync = false
	args := filter.Choose(strings.Split(cmd.Message.Text, " "), func(x string) bool {
		return x != ""
	}).([]string)
	name, args := args[0], args[1:]
	switch name {
	case "template":
		transitTo = func(m *t.Message) Handler {
			return SetTemplate{Message: m}
		}
		reply = "Жду шаблон объявления следующим сообщением."
	}
	return
}

type SetTemplate struct{ Message *t.Message }

func (cmd SetTemplate) Handle() (transitTo Transition, reply string, sync bool) {
	tpl := s.Template{ID: 0, SourceMessageID: 0, TargetMessageID: 0, Text: ""}
	s.Templates = append(s.Templates, tpl)
	transitTo = DefaultTransition
	reply = "Шаблон установлен"
	sync = true
	return
}
