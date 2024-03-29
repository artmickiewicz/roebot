package state

import (
	"bytes"
	"fmt"
	t "text/template"
)

type Chat struct {
	ID       int64
	Username string
	Title    string
}

type Template struct {
	ID               int
	TargetChannel    string
	SourceMessagePtr MessagePtr
	TargetMessagePtr MessagePtr
	Text             string
	TemplateObj      *t.Template
}

type MessagePtr struct {
	ChatID    int64
	MessageID int
}

func (tpl Template) IsPosted() bool {
	return tpl.TargetMessagePtr.MessageID > 0
}

func (tpl Template) Apply(vars map[string]string) string {
	if tpl.TemplateObj == nil {
		return ""
	}
	var buf bytes.Buffer
	if err := tpl.TemplateObj.Execute(&buf, vars); err != nil {
		return ""
	}
	return buf.String()
}

func (tpl Template) PrettyTarget() string {
	if tpl.TargetMessagePtr.MessageID > 0 {
		return fmt.Sprintf("%s#%d", tpl.TargetChannel, tpl.TargetMessagePtr.MessageID)
	} else {
		return tpl.TargetChannel
	}
}

type State int

const Null State = 0
const Clean State = 1
const Added State = 2
const Updated State = 3
const Deleted State = 4
